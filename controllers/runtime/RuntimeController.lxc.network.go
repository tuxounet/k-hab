package runtime

import (
	"fmt"

	"github.com/tuxounet/k-hab/utils"
)

func (r *RuntimeController) provisionNetwork() error {

	host_interface := r.ctx.GetConfigValue("hab.incus.host.interface")
	host_address := r.ctx.GetConfigValue("hab.incus.host.address")
	host_address_mask := r.ctx.GetConfigValue("hab.incus.host.netmask")
	network_nat := r.ctx.GetConfigValue("hab.incus.host.nat")

	ipv4_address := fmt.Sprintf("%s/%s", host_address, host_address_mask)
	existsNetwork, err := r.existsNetwork(host_interface)
	if err != nil {
		return err
	}

	if !existsNetwork {
		err = r.createNetwork(host_interface, "bridge", "ipv4.address="+ipv4_address, "ipv4.nat="+network_nat)
		if err != nil {
			return err
		}
	}
	return nil

}

func (r *RuntimeController) existsNetwork(name string) (bool, error) {

	cmd, err := r.withIncusCmd("network", "ls", "--format", "json")
	if err != nil {
		return false, err
	}
	out, err := utils.JsonCommandOutput[[]map[string]interface{}](cmd)
	if err != nil {
		return false, err
	}
	for _, profile := range out {
		if profile["name"] == name {
			return true, nil
		}
	}

	return false, nil

}

func (r *RuntimeController) createNetwork(name string, driver string, options ...string) error {

	cmd, err := r.withIncusCmd("network", "create", name, "--type", driver)
	if err != nil {
		return err
	}
	cmd.Args = append(cmd.Args, options...)

	return utils.OsExec(cmd)

}
