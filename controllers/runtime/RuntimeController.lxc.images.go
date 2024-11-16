package runtime

import (
	"github.com/tuxounet/k-hab/utils"
)

func (l *RuntimeController) PresentImage(name string) (bool, error) {
	cmd, err := l.withIncusCmd("image", "list", "--format", "json")
	if err != nil {
		return false, err
	}
	out, err := utils.JsonCommandOutput[[]map[string]interface{}](cmd)
	if err != nil {
		return false, err
	}

	for _, image := range out {
		aliases := image["aliases"].([]interface{})
		for _, ali := range aliases {
			alias := ali.(map[string]interface{})
			if alias["name"].(string) == name {
				return true, nil
			}
		}
	}
	return false, nil

}

func (l *RuntimeController) RegisterImage(name string, metadataPackage string, rootfsPackage string, force bool) error {

	present, err := l.PresentImage(name)
	if err != nil {
		return err
	}
	if present || force {
		err = l.RemoveImage(name)
		if err != nil {
			return err
		}

	}
	cmd, err := l.withIncusCmd("image", "import", metadataPackage, rootfsPackage, "--alias", name)
	if err != nil {
		return err
	}

	return utils.OsExec(cmd)

}

func (l *RuntimeController) RemoveImage(name string) error {

	present, err := l.PresentImage(name)
	if err != nil {
		return err
	}
	if present {
		cmd, err := l.withIncusCmd("image", "delete", name)
		if err != nil {
			return err
		}
		err = utils.OsExec(cmd)
		if err != nil {
			return err
		}

	}
	return nil

}
