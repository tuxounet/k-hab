package containers

import "errors"

func (r *ContainersController) loadContainers() error {

	for _, confContainer := range r.ctx.GetContainersConfig() {

		found := false
		for _, localContainer := range r.containers {
			if localContainer.name == confContainer.Name {
				found = true
				break
			}
		}
		if !found {
			container := NewContainerModel(confContainer.Name, r.ctx, confContainer)
			r.containers[container.name] = *container
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
		if container.name == name {
			return &container, nil
		}
	}

	return nil, errors.New("container not found")

}
