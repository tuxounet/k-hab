package images

import (
	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/controllers/builder"
	"github.com/tuxounet/k-hab/controllers/plateform"
)

func (c *ImageModel) getBuilderController() (*builder.BuilderController, error) {
	controller, err := c.ctx.GetController(bases.BuilderController)
	if err != nil {
		return nil, err
	}
	builderController := controller.(*builder.BuilderController)

	return builderController, nil
}

func (c *ImageModel) getPlateformController() (*plateform.PlateformController, error) {
	controller, err := c.ctx.GetController(bases.PlateformController)
	if err != nil {
		return nil, err
	}
	plateformController := controller.(*plateform.PlateformController)

	return plateformController, nil

}
