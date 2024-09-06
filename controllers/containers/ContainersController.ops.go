package containers

import "errors"

func (r *ContainersController) loadContainers() error {

	for _, confContainer := range r.ctx.GetSetupContainers() {

		found := false
		for _, localContainer := range r.containers {
			if localContainer.Name == confContainer.Name {
				found = true
				break
			}
		}
		if !found {
			containersPath, err := r.getContainersPath()
			if err != nil {
				return err
			}
			container := NewContainerModel(confContainer.Name, r.ctx, confContainer, containersPath)
			r.containers[container.Name] = *container
		}
	}

	return nil

}

func (r *ContainersController) GetContainer(name string) (*ContainerModel, error) {
	err := r.loadContainers()
	if err != nil {
		return nil, err
	}
	for _, container := range r.containers {
		if container.Name == name {
			return &container, nil
		}
	}

	return nil, errors.New("container not found")

}
