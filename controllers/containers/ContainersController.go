package containers

import (
	"github.com/tuxounet/k-hab/bases"
)

type ContainersController struct {
	bases.BaseController
	ctx        bases.IContext
	log        bases.ILogger
	containers map[string]ContainerModel
}

func NewContainersController(ctx bases.IContext) *ContainersController {

	return &ContainersController{
		ctx:        ctx,
		log:        ctx.GetSubLogger("ContainersController", ctx.GetLogger()),
		containers: make(map[string]ContainerModel),
	}
}
