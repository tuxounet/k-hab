package runtime

import (
	"fmt"

	"github.com/tuxounet/k-hab/utils"
)

func (r *RuntimeController) provisionNetwork() error {

	host_interface := utils.GetMapValue(r.ctx.GetHabConfig(), "lxd.lxc.host.interface").(string)
	host_address := utils.GetMapValue(r.ctx.GetHabConfig(), "lxd.lxc.host.address").(string)
	host_address_mask := utils.GetMapValue(r.ctx.GetHabConfig(), "lxd.lxc.host.netmask").(string)
	network_nat := utils.GetMapValue(r.ctx.GetHabConfig(), "lxd.lxc.host.nat").(string)

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

	cmd, err := r.withLxcCmd("network", "ls", "--format", "json")
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

	cmd, err := r.withLxcCmd("network", "create", name, "--type", driver)
	if err != nil {
		return err
	}
	cmd.Args = append(cmd.Args, options...)

	return utils.OsExec(cmd)

}
