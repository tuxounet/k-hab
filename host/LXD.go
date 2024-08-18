package host

import (
	"github.com/tuxounet/k-hab/utils"
)

type LXD struct {
	scopeBase string
	habConfig map[string]interface{}
}

func NewLXD(habConfig map[string]interface{}) *LXD {

	return &LXD{
		scopeBase: "LXD",
		habConfig: habConfig,
	}
}

func (l *LXD) withLxdCmd(ctx *utils.ScopeContext, args ...string) *utils.CmdCall {
	return utils.ScopingWithReturnOnly(ctx, l.scopeBase, "Present", func(ctx *utils.ScopeContext) *utils.CmdCall {
		return utils.WithCmdCallBuilder(ctx, l.habConfig, "lxd.command.prefix", "lxd.command.name", args...)
	})
}

func (l *LXD) withLxcCmd(ctx *utils.ScopeContext, args ...string) *utils.CmdCall {
	return utils.ScopingWithReturnOnly(ctx, l.scopeBase, "Present", func(ctx *utils.ScopeContext) *utils.CmdCall {
		return utils.WithCmdCallBuilder(ctx, l.habConfig, "lxd.lxc.command.prefix", "lxd.lxc.command.name", args...)
	})
}

func (l *LXD) Present(ctx *utils.ScopeContext) bool {
	return utils.ScopingWithReturnOnly(ctx, l.scopeBase, "Present", func(ctx *utils.ScopeContext) bool {
		snaps := NewSnapPackages(l.habConfig)

		snapName := utils.GetMapValue(ctx, l.habConfig, "lxd.snap").(string)

		present := snaps.InstalledSnap(ctx, snapName)

		return present
	})
}

func (l *LXD) Provision(ctx *utils.ScopeContext) error {

	return ctx.Scope(l.scopeBase, "Provision", func(ctx *utils.ScopeContext) {

		snaps := NewSnapPackages(l.habConfig)
		snapName := utils.GetMapValue(ctx, l.habConfig, "lxd.snap").(string)
		present := snaps.InstalledSnap(ctx, snapName)

		if !present {
			snapMode := utils.GetMapValue(ctx, l.habConfig, "lxd.snap_mode").(string)
			ctx.Must(snaps.InstallSnap(ctx, snapName, snapMode))
		}
		ctx.Must(utils.ExecSyncOutput(ctx, l.withLxdCmd(ctx, "waitready")))
		ctx.Must(l.ProvisionStorage(ctx))
		ctx.Must(l.ProvisionNetwork(ctx))
		ctx.Must(l.ProvisionProfile(ctx))
	})

}

func (l *LXD) Up(ctx *utils.ScopeContext) error {
	return ctx.Scope(l.scopeBase, "Up", func(ctx *utils.ScopeContext) {
		ctx.Log.InfoF("READY")
	})
}

func (l *LXD) Down(ctx *utils.ScopeContext) error {
	return ctx.Scope(l.scopeBase, "Down", func(ctx *utils.ScopeContext) {
		snaps := NewSnapPackages(l.habConfig)
		snapName := utils.GetMapValue(ctx, l.habConfig, "lxd.snap").(string)
		present := snaps.InstalledSnap(ctx, snapName)

		if !present {
			return
		}

		ctx.Must(utils.ExecSyncOutput(ctx, l.withLxdCmd(ctx, "shutdown")))
	})

}
func (l *LXD) Unprovision(ctx *utils.ScopeContext) error {
	return ctx.Scope(l.scopeBase, "Unprovision", func(ctx *utils.ScopeContext) {
		snaps := NewSnapPackages(l.habConfig)
		snapName := utils.GetMapValue(ctx, l.habConfig, "lxd.snap").(string)
		present := snaps.InstalledSnap(ctx, snapName)

		if !present {
			return
		}

		ctx.Must(l.UnprovisionProfile(ctx))
		ctx.Must(l.UnprovisionStorage(ctx))
	})

}
func (l *LXD) Nuke(ctx *utils.ScopeContext) error {
	return ctx.Scope(l.scopeBase, "Nuke", func(ctx *utils.ScopeContext) {
		ctx.Must(l.NukeStorage(ctx))
		snaps := NewSnapPackages(l.habConfig)
		snapName := utils.GetMapValue(ctx, l.habConfig, "lxd.snap").(string)
		ctx.Must(snaps.RemoveSnap(ctx, snapName))
		ctx.Must(snaps.RemoveSnapSnapshots(ctx, snapName))
	})
}
