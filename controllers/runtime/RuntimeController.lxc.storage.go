package runtime

import (
	"os"
	"path"

	"github.com/tuxounet/k-hab/utils"
)

func (r *RuntimeController) getStoragePath() (string, error) {

	storageRoot, err := r.ctx.GetStorageRoot()
	if err != nil {
		return "", err
	}
	storagePathDefinition := r.ctx.GetConfigValue("hab.lxd.lxc.storage.path")
	storagePath := path.Join(storageRoot, storagePathDefinition)

	err = os.MkdirAll(storagePath, 0755)
	if err != nil {
		return "", err
	}
	return storagePath, nil

}

func (r *RuntimeController) provisionStorage() error {

	storage_pool := r.ctx.GetConfigValue("hab.lxd.lxc.storage.pool")
	storage_driver := r.ctx.GetConfigValue("hab.lxd.lxc.storage.driver")
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

func (r *RuntimeController) unprovisionStorage() error {

	storage_pool := r.ctx.GetConfigValue("hab.lxd.lxc.storage.pool")

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

func (r *RuntimeController) nukeStorage() error {

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

func (r *RuntimeController) existsStorage(name string) (bool, error) {

	cmd, err := r.withLxcCmd("storage", "ls", "--format", "json")
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

func (r *RuntimeController) createStorage(name string, driver string, options ...string) error {

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

	cmd, err := r.withLxcCmd("storage", "create", name, driver)
	if err != nil {
		return err
	}

	cmd.Args = append(cmd.Args, options...)

	return utils.OsExec(cmd)

}

func (r *RuntimeController) removeStorage(name string) error {

	cmd, err := r.withLxcCmd("storage", "delete", name)
	if err != nil {
		return err
	}
	return utils.OsExec(cmd)

}
