package context

import (
	"errors"
	"os"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/controllers/builder"
	"github.com/tuxounet/k-hab/controllers/containers"
	"github.com/tuxounet/k-hab/controllers/dependencies"
	"github.com/tuxounet/k-hab/controllers/egress"
	"github.com/tuxounet/k-hab/controllers/images"
	"github.com/tuxounet/k-hab/controllers/ingress"
	"github.com/tuxounet/k-hab/controllers/runtime"
)

func (h *HabContext) Init() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	h.cwd = cwd

	err = h.config.Load()
	if err != nil {
		return err
	}
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
		case bases.RuntimeController:
			controller = runtime.NewRuntimeController(h)
		case bases.ContainersController:
			controller = containers.NewContainersController(h)
		case bases.EgressController:
			controller = egress.NewHttpEgressController(h)
		case bases.ImagesController:
			controller = images.NewImagesController(h)
		default:
			return errors.New("invalid controller name " + string(controllerKey))
		}

		if controller == nil {
			return errors.New("iontroller is nil")
		}

		h.controllers[bases.HabControllers(controllerKey)] = controller
	}

	return nil
}
