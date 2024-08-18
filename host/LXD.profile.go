package host

import (
	"strings"

	"github.com/tuxounet/k-hab/utils"
)

func (l *LXD) ProvisionProfile(ctx *utils.ScopeContext) error {
	return ctx.Scope(l.scopeBase, "ProvisionProfile", func(ctx *utils.ScopeContext) {

		profile := utils.GetMapValue(ctx, l.habConfig, "lxd.lxc.profile").(string)
		storage_pool := utils.GetMapValue(ctx, l.habConfig, "lxd.lxc.storage.pool").(string)
		network_bridge := utils.GetMapValue(ctx, l.habConfig, "lxd.lxc.host.interface").(string)

		profileExists := l.existsProfile(ctx, profile)
		if !profileExists {
			ctx.Must(l.createProfile(ctx, profile))
		}

		networkDeviceExists := l.existsDeviceProfile(ctx, profile, "eth0")
		if !networkDeviceExists {
			//lxc --debug  profile device add plop eth0 nic  nictype=bridged parent=lxbr0 name=eth0
			ctx.Must(l.addDeviceProfile(ctx, profile, "eth0", "nic", "nictype=bridged", "parent="+network_bridge, "name=eth0"))
		}

		rootStorageExists := l.existsDeviceProfile(ctx, profile, "root")
		if !rootStorageExists {
			ctx.Must(l.addDeviceProfile(ctx, profile, "root", "disk", "path=/", "pool="+storage_pool))
		}
	})
}

func (l *LXD) UnprovisionProfile(ctx *utils.ScopeContext) error {
	return ctx.Scope(l.scopeBase, "UnprovisionProfile", func(ctx *utils.ScopeContext) {
		profile := utils.GetMapValue(ctx, l.habConfig, "lxd.lxc.profile").(string)
		profileExists := l.existsProfile(ctx, profile)
		if profileExists {
			ctx.Must(l.deleteProfile(ctx, profile))
		}
	})
}

func (l *LXD) existsProfile(ctx *utils.ScopeContext, name string) bool {

	return utils.ScopingWithReturn(ctx, l.scopeBase, "existsProfile", func(ctx *utils.ScopeContext) bool {

		arr := utils.CommandSyncJsonArrayOutput(ctx, l.withLxcCmd(ctx, "profile", "ls", "--format", "json"))

		for _, profile := range arr {
			if profile["name"] == name {
				return true
			}
		}

		return false
	})
}

func (l *LXD) createProfile(ctx *utils.ScopeContext, name string) error {
	return ctx.Scope(l.scopeBase, "createProfile", func(ctx *utils.ScopeContext) {
		ctx.Must(utils.ExecSyncOutput(ctx, l.withLxcCmd(ctx, "profile", "create", name)))
	})
}
func (l *LXD) deleteProfile(ctx *utils.ScopeContext, name string) error {
	return ctx.Scope(l.scopeBase, "deleteProfile", func(ctx *utils.ScopeContext) {
		ctx.Must(utils.ExecSyncOutput(ctx, l.withLxcCmd(ctx, "profile", "delete", name)))
	})
}

func (l *LXD) existsDeviceProfile(ctx *utils.ScopeContext, profileName string, deviceName string) bool {

	return utils.ScopingWithReturn(ctx, l.scopeBase, "existsDeviceProfile", func(ctx *utils.ScopeContext) bool {

		out := utils.CommandSyncOutput(ctx, l.withLxcCmd(ctx, "profile", "device", "list", profileName))

		lines := strings.Split(out, "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, deviceName) {
				return true
			}
		}

		return false
	})
}

func (l *LXD) addDeviceProfile(ctx *utils.ScopeContext, profileName string, deviceName string, deviceType string, options ...string) error {
	return ctx.Scope(l.scopeBase, "AddDeviceProfile", func(ctx *utils.ScopeContext) {

		cmd := l.withLxcCmd(ctx, "profile", "device", "add", profileName, deviceName, deviceType)
		cmd.Args = append(cmd.Args, options...)

		ctx.Must(utils.ExecSyncOutput(ctx, cmd))
	})
}
