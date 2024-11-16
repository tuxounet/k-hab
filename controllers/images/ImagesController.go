package images

import (
	"github.com/tuxounet/k-hab/bases"
)

type ImagesController struct {
	bases.BaseController
	ctx    bases.IContext
	log    bases.ILogger
	images []*ImageModel
}

func NewImagesController(ctx bases.IContext) *ImagesController {
	return &ImagesController{
		ctx: ctx,
		log: ctx.GetSubLogger(string(bases.ImagesController), ctx.GetLogger()),
	}
}
