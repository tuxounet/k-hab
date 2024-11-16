package plateform

import (
	"github.com/tuxounet/k-hab/bases"
)

type PlateformController struct {
	bases.BaseController
	ctx bases.IContext
	log bases.ILogger
}

func NewPlateformController(ctx bases.IContext) *PlateformController {
	return &PlateformController{
		ctx: ctx,
		log: ctx.GetSubLogger(string(bases.PlateformController), ctx.GetLogger()),
	}
}

func (c *PlateformController) IsPresent() (bool, error) {
	return c.presentService()
}
