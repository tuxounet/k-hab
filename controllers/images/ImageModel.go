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
	Name       string
	Definition definitions.HabBaseDefinition
}

func NewImageModel(name string, ctx bases.IContext, definition definitions.HabBaseDefinition) *ImageModel {

	return &ImageModel{
		Name:       name,
		ctx:        ctx,
		Definition: definition,
	}
}

func (hi *ImageModel) present() (bool, error) {
	rRunContaner, err := hi.ctx.GetController("RuntimeController")
	if err != nil {
		return false, err
	}
	runtimeController := rRunContaner.(*runtime.RuntimeController)

	return runtimeController.PresentImage(hi.Name)

}

func (hi *ImageModel) needBuild(definition definitions.HabBaseDefinition) (bool, error) {
	rBuildContainer, err := hi.ctx.GetController("BuilderController")
	if err != nil {
		return false, err
	}
	builderController := rBuildContainer.(*builder.BuilderController)

	sExpectedBuilderConfig, err := utils.UnTemplate(definition.Builder, map[string]interface{}{
		"config": hi.ctx.GetCurrentConfig(),
		"image":  hi.Definition,
	})
	if err != nil {
		return false, err
	}

	return builderController.ConfigHasChanged(hi.Name, sExpectedBuilderConfig)

}

func (hi *ImageModel) provision() error {

	sBuilderConfig, err := utils.UnTemplate(hi.Definition.Builder, map[string]interface{}{
		"config": hi.ctx.GetCurrentConfig(),
		"image":  hi.Definition,
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

	buildResult, err := builderController.BuildDistro(hi.Name, sBuilderConfig)
	if err != nil {
		return err
	}

	err = runtimeController.RegisterImage(hi.Name, buildResult.MetadataPackage, buildResult.RootfsPackage, buildResult.Built)
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

	rRunContaner, err := hi.ctx.GetController("RuntimeController")
	if err != nil {
		return err
	}
	runtimeController := rRunContaner.(*runtime.RuntimeController)

	err = runtimeController.RemoveImage(hi.Name)
	if err != nil {
		return err
	}

	err = builderController.RemoveCache(hi.Name)
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
	err = builderController.RemoveCache(hi.Name)
	if err != nil {
		return err

	}
	return nil
}
