package context

import (
	"errors"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/controllers/containers"
)

func (h *HabContext) getEntryContainer() (*containers.ContainerModel, error) {
	setupContainers := h.GetSetupContainers()
	if len(setupContainers) == 0 {
		return nil, errors.New("no container found in setup")
	}
	entryContainer := setupContainers[0]
	entrypoint := entryContainer.Name

	containersController, err := h.GetController(bases.ContainersController)

	if err != nil {
		return nil, err
	}

	containers := containersController.(*containers.ContainersController)

	container, err := containers.GetContainer(entrypoint)

	if err != nil {
		return nil, err
	}

	return container, nil

}
