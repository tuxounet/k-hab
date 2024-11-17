package plateform

import (
	"fmt"
	"strings"

	"github.com/tuxounet/k-hab/utils"
)

func (r *PlateformController) provisionProfile() error {

	profile := r.ctx.GetConfigValue("hab.name")

	if_name := fmt.Sprintf("%s0", profile)

	profileExists, err := r.existsProfile(profile)
	if err != nil {
		return err
	}
	if !profileExists {
		err = r.createProfile(profile)
		if err != nil {
			return err
		}
	}

	networkDeviceExists, err := r.existsDeviceProfile(profile, "eth0")
	if err != nil {
		return err
	}
	if !networkDeviceExists {
		err = r.addDeviceProfile(profile, "eth0", "nic", "nictype=bridged", "parent="+if_name, "name=eth0")
		if err != nil {
			return err
		}
	}

	rootStorageExists, err := r.existsDeviceProfile(profile, "root")
	if err != nil {
		return err
	}
	if !rootStorageExists {
		err = r.addDeviceProfile(profile, "root", "disk", "path=/", "pool="+profile)
		if err != nil {
			return err
		}
	}
	return nil

}

func (r *PlateformController) unprovisionProfile() error {

	profile := r.ctx.GetConfigValue("hab.name")
	profileExists, err := r.existsProfile(profile)
	if err != nil {
		return err
	}
	if profileExists {
		err = r.deleteProfile(profile)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PlateformController) existsProfile(name string) (bool, error) {

	cmd, err := r.withLxcCmd("profile", "ls", "--format", "json")
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

func (r *PlateformController) createProfile(name string) error {
	cmd, err := r.withLxcCmd("profile", "create", name)
	if err != nil {
		return err
	}
	return utils.OsExec(cmd)

}
func (r *PlateformController) deleteProfile(name string) error {
	cmd, err := r.withLxcCmd("profile", "delete", name)

	if err != nil {
		return err
	}

	return utils.OsExec(cmd)

}

func (r *PlateformController) existsDeviceProfile(profileName string, deviceName string) (bool, error) {

	cmd, err := r.withLxcCmd("profile", "device", "list", profileName)

	if err != nil {
		return false, err
	}
	out, err := utils.RawCommandOutput(cmd)

	if err != nil {
		return false, err
	}

	lines := strings.Split(out, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, deviceName) {
			return true, nil
		}
	}

	return false, nil

}

func (r *PlateformController) addDeviceProfile(profileName string, deviceName string, deviceType string, options ...string) error {

	cmd, err := r.withLxcCmd("profile", "device", "add", profileName, deviceName, deviceType)
	if err != nil {
		return err
	}
	cmd.Args = append(cmd.Args, options...)

	return utils.OsExec(cmd)

}
