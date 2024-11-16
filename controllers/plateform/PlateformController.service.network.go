package plateform

import (
	"fmt"

	"github.com/tuxounet/k-hab/utils"
)

func (r *PlateformController) provisionNetwork() error {

	host_interface := r.ctx.GetConfigValue("hab.incus.host.interface")
	host_v4_address := r.ctx.GetConfigValue("hab.incus.host.v4.address")
	host_v4_address_mask := r.ctx.GetConfigValue("hab.incus.host.v4.netmask")
	host_v4_nat := r.ctx.GetConfigValue("hab.incus.host.v4.nat")
	host_v6_address := r.ctx.GetConfigValue("hab.incus.host.v6.address")
	host_v6_nat := r.ctx.GetConfigValue("hab.incus.host.v6.nat")

	ipv4_address := fmt.Sprintf("%s/%s", host_v4_address, host_v4_address_mask)
	existsNetwork, err := r.existsNetwork(host_interface)
	if err != nil {
		return err
	}

	if !existsNetwork {
		err = r.createNetwork(host_interface, "bridge", "ipv4.address="+ipv4_address, "ipv4.nat="+host_v4_nat, "ipv6.address="+host_v6_address, "ipv6.nat="+host_v6_nat)
		if err != nil {
			return err
		}
	}
	return nil

}

func (r *PlateformController) unprovisionNetwork() error {

	host_interface := r.ctx.GetConfigValue("hab.incus.host.interface")
	existsNetwork, err := r.existsNetwork(host_interface)
	if err != nil {
		return err
	}
	if existsNetwork {
		err = r.deleteNetwork(host_interface)
		if err != nil {
			return err
		}
	}
	return nil

}

func (r *PlateformController) existsNetwork(name string) (bool, error) {

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

func (r *PlateformController) createNetwork(name string, driver string, options ...string) error {

	cmd, err := r.withIncusCmd("network", "create", name, "--type", driver)
	if err != nil {
		return err
	}
	cmd.Args = append(cmd.Args, options...)

	return utils.OsExec(cmd)

}

func (r *PlateformController) deleteNetwork(name string) error {

	cmd, err := r.withIncusCmd("network", "delete", name)

	if err != nil {
		return err
	}
	return utils.OsExec(cmd)

}
