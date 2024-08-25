package runtime

import (
	"github.com/tuxounet/k-hab/bases"
)

type RuntimeController struct {
	bases.BaseController
	ctx bases.IContext
	log bases.ILogger
}

func NewRuntimeController(ctx bases.IContext) *RuntimeController {
	return &RuntimeController{
		ctx: ctx,
		log: ctx.GetSubLogger("RuntimeController", ctx.GetLogger()),
	}
}
