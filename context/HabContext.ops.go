package context

import (
	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/controllers/containers"
	"github.com/tuxounet/k-hab/utils"
)

func (h *HabContext) getEntryContainer() (*containers.ContainerModel, error) {

	entrypoint := utils.GetMapValue(h.GetHabConfig(), "entry.container").(string)

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
