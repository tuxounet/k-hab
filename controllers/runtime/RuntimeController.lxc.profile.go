package runtime

import (
	"strings"

	"github.com/tuxounet/k-hab/utils"
)

func (r *RuntimeController) provisionProfile() error {

	profile := r.ctx.GetConfigValue("hab.incus.profile")
	storage_pool := r.ctx.GetConfigValue("hab.incus.storage.pool")
	network_bridge := r.ctx.GetConfigValue("hab.incus.host.interface")

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
		//lxc --debug  profile device add plop eth0 nic  nictype=bridged parent=lxbr0 name=eth0
		err = r.addDeviceProfile(profile, "eth0", "nic", "nictype=bridged", "parent="+network_bridge, "name=eth0")
		if err != nil {
			return err
		}
	}

	rootStorageExists, err := r.existsDeviceProfile(profile, "root")
	if err != nil {
		return err
	}
	if !rootStorageExists {
		err = r.addDeviceProfile(profile, "root", "disk", "path=/", "pool="+storage_pool)
		if err != nil {
			return err
		}
	}
	return nil

}

func (r *RuntimeController) unprovisionProfile() error {

	profile := r.ctx.GetConfigValue("hab.incus.profile")
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

func (r *RuntimeController) existsProfile(name string) (bool, error) {

	cmd, err := r.withIncusCmd("profile", "ls", "--format", "json")
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

func (r *RuntimeController) createProfile(name string) error {
	cmd, err := r.withIncusCmd("profile", "create", name)
	if err != nil {
		return err
	}
	return utils.OsExec(cmd)

}
func (r *RuntimeController) deleteProfile(name string) error {
	cmd, err := r.withIncusCmd("profile", "delete", name)

	if err != nil {
		return err
	}

	return utils.OsExec(cmd)

}

func (r *RuntimeController) existsDeviceProfile(profileName string, deviceName string) (bool, error) {

	cmd, err := r.withIncusCmd("profile", "device", "list", profileName)

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

func (r *RuntimeController) addDeviceProfile(profileName string, deviceName string, deviceType string, options ...string) error {

	cmd, err := r.withIncusCmd("profile", "device", "add", profileName, deviceName, deviceType)
	if err != nil {
		return err
	}
	cmd.Args = append(cmd.Args, options...)

	return utils.OsExec(cmd)

}
