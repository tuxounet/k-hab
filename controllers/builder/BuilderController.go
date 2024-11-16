package builder

import (
	"github.com/tuxounet/k-hab/bases"
)

type BuilderController struct {
	bases.BaseController
	ctx bases.IContext
	log bases.ILogger
}

func NewBuilderController(ctx bases.IContext) *BuilderController {
	return &BuilderController{
		ctx: ctx,
		log: ctx.GetSubLogger(string(bases.BuilderController), ctx.GetLogger()),
	}
}
