package pki

import (
	"github.com/tuxounet/k-hab/bases"
)

type PKIController struct {
	bases.BaseController
	ctx bases.IContext
	log bases.ILogger
}

func NewPKIController(ctx bases.IContext) *PKIController {
	return &PKIController{
		ctx: ctx,
		log: ctx.GetSubLogger(string(bases.PKIController), ctx.GetLogger()),
	}
}
