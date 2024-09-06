package containers

import (
	"os"
	"path"
)

func (b *ContainersController) getContainersPath() (string, error) {

	storageRoot, err := b.ctx.GetStorageRoot()
	if err != nil {
		return "", err
	}
	containersPathDefinition := b.ctx.GetConfigValue("hab.containers.path")
	containersPath := path.Join(storageRoot, containersPathDefinition)

	err = os.MkdirAll(containersPath, 0755)
	if err != nil {
		return "", err
	}

	return containersPath, nil

}
