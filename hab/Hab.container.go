package hab

import (
	"strings"

	"github.com/tuxounet/k-hab/host"
	"github.com/tuxounet/k-hab/utils"
)

type HabContainer struct {
	scopeBase string
	name      string
	hab       *Hab
}

func newHabContainer(name string, hab *Hab) *HabContainer {
	base := "HabContainer/" + name
	return &HabContainer{
		scopeBase: base,
		name:      name,
		hab:       hab,
	}
}
func (hc *HabContainer) getLxc(ctx *utils.ScopeContext) *host.LXC {
	return utils.ScopingWithReturn(ctx, hc.scopeBase, "getLxc", func(ctx *utils.ScopeContext) *host.LXC {

		return host.NewLXC(hc.name, hc.hab.config.HabConfig, hc.hab.config.GetContainerConfig(ctx, hc.name))

	})
}

func (hc *HabContainer) provision(ctx *utils.ScopeContext) error {
	return ctx.Scope(hc.scopeBase, "provision", func(ctx *utils.ScopeContext) {
		ctx.Must(hc.getLxc(ctx).Provision(ctx))
	})
}

func (hc *HabContainer) up(ctx *utils.ScopeContext) error {
	return ctx.Scope(hc.scopeBase, "up", func(ctx *utils.ScopeContext) {
		ctx.Must(hc.getLxc(ctx).Start(ctx))
	})
}

func (hc *HabContainer) entry(ctx *utils.ScopeContext) error {
	return ctx.Scope(hc.scopeBase, "entry", func(ctx *utils.ScopeContext) {
		conf := hc.hab.config.GetContainerConfig(ctx, hc.name)
		shell_cmd := utils.GetMapValue(ctx, conf.ToMap(), "entry").(string)

		call := []string{shell_cmd}
		if strings.Contains(shell_cmd, " ") {
			call = strings.Split(shell_cmd, " ")
		}
		ctx.Must(hc.getLxc(ctx).Exec(ctx, call...))
	})
}

func (hc *HabContainer) shell(ctx *utils.ScopeContext) error {
	return ctx.Scope(hc.scopeBase, "shell", func(ctx *utils.ScopeContext) {
		conf := hc.hab.config.GetContainerConfig(ctx, hc.name)
		shell_cmd := utils.GetMapValue(ctx, conf.ToMap(), "shell").(string)

		call := []string{shell_cmd}
		if strings.Contains(shell_cmd, " ") {
			call = strings.Split(shell_cmd, " ")
		}
		ctx.Must(hc.getLxc(ctx).Exec(ctx, call...))
	})
}

func (hc *HabContainer) waitReady(ctx *utils.ScopeContext) error {
	return ctx.Scope(hc.scopeBase, "waitReady", func(ctx *utils.ScopeContext) {
		ctx.Must(hc.getLxc(ctx).WaitReady(ctx))
	})
}

func (hc *HabContainer) down(ctx *utils.ScopeContext) error {
	return ctx.Scope(hc.scopeBase, "down", func(ctx *utils.ScopeContext) {
		ctx.Must(hc.getLxc(ctx).Stop(ctx))
	})
}

func (hc *HabContainer) unprovision(ctx *utils.ScopeContext) error {
	return ctx.Scope(hc.scopeBase, "unprovision", func(ctx *utils.ScopeContext) {

		ctx.Must(hc.getLxc(ctx).Unprovision(ctx))

	})
}
