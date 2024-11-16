package containers

import (
	"os"
	"path"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/controllers/plateform"
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

func (c *ContainersController) getPlateformController() (*plateform.PlateformController, error) {
	controller, err := c.ctx.GetController(bases.PlateformController)
	if err != nil {
		return nil, err
	}
	plateformController := controller.(*plateform.PlateformController)
	return plateformController, nil

}
