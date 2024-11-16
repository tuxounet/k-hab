package plateform

import (
	"os"
	"path"

	"github.com/tuxounet/k-hab/utils"
)

func (r *PlateformController) getStoragePath() (string, error) {

	storageRoot, err := r.ctx.GetStorageRoot()
	if err != nil {
		return "", err
	}
	storagePathDefinition := r.ctx.GetConfigValue("hab.incus.storage.path")
	storagePath := path.Join(storageRoot, storagePathDefinition)

	err = os.MkdirAll(storagePath, 0755)
	if err != nil {
		return "", err
	}
	return storagePath, nil

}

func (r *PlateformController) provisionStorage() error {

	storage_pool := r.ctx.GetConfigValue("hab.incus.storage.pool")
	storage_driver := r.ctx.GetConfigValue("hab.incus.storage.driver")
	storage_path, err := r.getStoragePath()
	if err != nil {
		return err
	}

	storageExists, err := r.existsStorage(storage_pool)
	if err != nil {
		return err
	}

	if !storageExists {
		err = r.createStorage(storage_pool, storage_driver, "source="+storage_path)
		if err != nil {
			return err
		}

	}
	return nil
}

func (r *PlateformController) unprovisionStorage() error {

	storage_pool := r.ctx.GetConfigValue("hab.incus.storage.pool")

	storageExists, err := r.existsStorage(storage_pool)
	if err != nil {
		return err
	}
	if storageExists {
		err = r.removeStorage(storage_pool)
		if err != nil {

			return err
		}
	}
	return nil

}

func (r *PlateformController) nukeStorage() error {

	stroagePath, err := r.getStoragePath()
	if err != nil {
		return err
	}

	cmd, err := utils.WithCmdCall(r.ctx, "hab.rm.prefix", "hab.rm.name", "-rf", stroagePath)
	if err != nil {
		return err
	}

	return utils.OsExec(cmd)

}

func (r *PlateformController) existsStorage(name string) (bool, error) {

	cmd, err := r.withIncusCmd("storage", "ls", "--format", "json")
	if err != nil {
		return false, err
	}
	arr, err := utils.JsonCommandOutput[[]map[string]interface{}](cmd)
	if err != nil {
		return false, err
	}

	for _, profile := range arr {
		if profile["name"] == name {
			return true, nil
		}
	}

	return false, nil

}

func (r *PlateformController) createStorage(name string, driver string, options ...string) error {

	err := r.nukeStorage()
	if err != nil {
		return err
	}
	storage_path, err := r.getStoragePath()

	if err != nil {
		return err
	}

	err = os.MkdirAll(storage_path, 0755)
	if err != nil {
		return err
	}

	cmd, err := r.withIncusCmd("storage", "create", name, driver)
	if err != nil {
		return err
	}

	cmd.Args = append(cmd.Args, options...)

	return utils.OsExec(cmd)

}

func (r *PlateformController) removeStorage(name string) error {

	cmd, err := r.withIncusCmd("storage", "delete", name)
	if err != nil {
		return err
	}
	return utils.OsExec(cmd)

}