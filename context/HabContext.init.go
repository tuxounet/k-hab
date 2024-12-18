package context

import (
	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/controllers/builder"
	"github.com/tuxounet/k-hab/controllers/containers"
	"github.com/tuxounet/k-hab/controllers/dependencies"
	"github.com/tuxounet/k-hab/controllers/egress"
	"github.com/tuxounet/k-hab/controllers/images"
	"github.com/tuxounet/k-hab/controllers/ingress"
	"github.com/tuxounet/k-hab/controllers/pki"
	"github.com/tuxounet/k-hab/controllers/plateform"
)

func (h *HabContext) Init() error {

	order := bases.HabControllersLoadOrder()
	h.controllers = make(map[bases.HabControllers]bases.IController, len(order))

	for _, controllerKey := range order {
		var controller bases.IController

		switch controllerKey {
		case bases.DependenciesController:
			controller = dependencies.NewDependenciesController(h)
		case bases.BuilderController:
			controller = builder.NewBuilderController(h)
		case bases.IngressController:
			controller = ingress.NewHttpIngress(h)
		case bases.PlateformController:
			controller = plateform.NewPlateformController(h)
		case bases.ContainersController:
			controller = containers.NewContainersController(h)
		case bases.EgressController:
			controller = egress.NewHttpEgressController(h)
		case bases.ImagesController:
			controller = images.NewImagesController(h)
		case bases.PKIController:
			controller = pki.NewPKIController(h)
		}

		h.controllers[bases.HabControllers(controllerKey)] = controller
	}

	return nil
}
