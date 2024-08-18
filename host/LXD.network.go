package host

import (
	"fmt"

	"github.com/tuxounet/k-hab/utils"
)

func (l *LXD) ProvisionNetwork(ctx *utils.ScopeContext) error {
	return ctx.Scope(l.scopeBase, "ProvisionNetwork", func(ctx *utils.ScopeContext) {

		host_interface := utils.GetMapValue(ctx, l.habConfig, "lxd.lxc.host.interface").(string)
		host_address := utils.GetMapValue(ctx, l.habConfig, "lxd.lxc.host.address").(string)
		host_address_mask := utils.GetMapValue(ctx, l.habConfig, "lxd.lxc.host.netmask").(string)
		network_nat := utils.GetMapValue(ctx, l.habConfig, "lxd.lxc.host.nat").(string)

		ipv4_address := fmt.Sprintf("%s/%s", host_address, host_address_mask)
		existsNetwork := l.existsNetwork(ctx, host_interface)

		if !existsNetwork {
			ctx.Must(l.createNetwork(ctx, host_interface, "bridge", "ipv4.address="+ipv4_address, "ipv4.nat="+network_nat))
		}
	})

}

func (l *LXD) existsNetwork(ctx *utils.ScopeContext, name string) bool {

	return utils.ScopingWithReturn(ctx, l.scopeBase, "existsNetwork", func(ctx *utils.ScopeContext) bool {

		arr := utils.CommandSyncJsonArrayOutput(ctx, l.withLxcCmd(ctx, "network", "ls", "--format", "json"))
		for _, profile := range arr {
			if profile["name"] == name {
				return true
			}
		}

		return false
	})
}

func (l *LXD) createNetwork(ctx *utils.ScopeContext, name string, driver string, options ...string) error {
	return ctx.Scope(l.scopeBase, "createNetwork", func(ctx *utils.ScopeContext) {

		cmd := l.withLxcCmd(ctx, "network", "create", name, "--type", driver)
		cmd.Args = append(cmd.Args, options...)

		ctx.Must(utils.ExecSyncOutput(ctx, cmd))
	})

}
