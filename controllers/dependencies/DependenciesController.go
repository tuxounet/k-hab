package dependencies

import (
	"github.com/tuxounet/k-hab/bases"
)

type DependenciesController struct {
	bases.BaseController
	ctx bases.IContext
	log bases.ILogger
}

func NewDependenciesController(ctx bases.IContext) *DependenciesController {
	return &DependenciesController{
		ctx: ctx,
		log: ctx.GetSubLogger("DependenciesController", ctx.GetLogger()),
	}
}
