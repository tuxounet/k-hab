package images

import (
	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/controllers/builder"
	"github.com/tuxounet/k-hab/controllers/images/definitions"
	"github.com/tuxounet/k-hab/controllers/runtime"
	"github.com/tuxounet/k-hab/utils"
)

type ImageModel struct {
	ctx        bases.IContext
	name       string
	definition definitions.HabBaseDefinition
}

func NewImageModel(name string, ctx bases.IContext, definition definitions.HabBaseDefinition) *ImageModel {

	return &ImageModel{
		name:       name,
		ctx:        ctx,
		definition: definition,
	}
}

func (hi *ImageModel) present() (bool, error) {
	rRunContaner, err := hi.ctx.GetController("RuntimeController")
	if err != nil {
		return false, err
	}
	runtimeController := rRunContaner.(*runtime.RuntimeController)

	return runtimeController.PresentImage(hi.name)

}

func (hi *ImageModel) provision() error {
	habConfig := hi.ctx.GetHabConfig()

	sBuilderConfig, err := utils.UnTemplate(hi.definition.Builder, map[string]interface{}{
		"hab":   habConfig,
		"image": hi.definition,
	})
	if err != nil {
		return err
	}

	rBuildContainer, err := hi.ctx.GetController("BuilderController")
	if err != nil {
		return err
	}
	builderController := rBuildContainer.(*builder.BuilderController)

	rRunContaner, err := hi.ctx.GetController("RuntimeController")
	if err != nil {
		return err
	}
	runtimeController := rRunContaner.(*runtime.RuntimeController)

	buildResult, err := builderController.BuildDistro(hi.name, sBuilderConfig)
	if err != nil {
		return err
	}

	err = runtimeController.RegisterImage(hi.name, buildResult.MetadataPackage, buildResult.RootfsPackage, buildResult.Built)
	if err != nil {
		return err
	}
	return nil

}

func (hi *ImageModel) unprovision() error {

	controller, err := hi.ctx.GetController("BuilderController")
	if err != nil {
		return err
	}
	builderController := controller.(*builder.BuilderController)
	err = builderController.RemoveCache(hi.name)
	if err != nil {
		return err
	}
	return nil

}

func (hi *ImageModel) nuke() error {

	controller, err := hi.ctx.GetController("BuilderController")
	if err != nil {
		return err
	}
	builderController := controller.(*builder.BuilderController)
	err = builderController.RemoveCache(hi.name)
	if err != nil {
		return err

	}
	return nil
}
