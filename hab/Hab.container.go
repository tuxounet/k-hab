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
	lxc       *host.LXC
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
	return utils.ScopingWithReturnOnly(ctx, hc.scopeBase, "getLxc", func(ctx *utils.ScopeContext) *host.LXC {

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

func (hc *HabContainer) shell(ctx *utils.ScopeContext) error {
	return ctx.Scope(hc.scopeBase, "shell", func(ctx *utils.ScopeContext) {
		shell_cmd := utils.GetMapValue(ctx, hc.getLxc(ctx).ContainerConfig, "shell").(string)

		call := make([]string, 0)
		if strings.Contains(shell_cmd, " ") {
			call = strings.Split(shell_cmd, " ")
		} else {
			call = []string{shell_cmd}
		}

		ctx.Must(hc.getLxc(ctx).Exec(ctx, call...))
	})
}

func (hc *HabContainer) waitReady(ctx *utils.ScopeContext) error {
	return ctx.Scope(hc.scopeBase, "waitReady", func(ctx *utils.ScopeContext) {
		ctx.Must(hc.getLxc(ctx).WaitReady(ctx))
	})
}

func (hc *HabContainer) exec(ctx *utils.ScopeContext) error {
	return ctx.Scope(hc.scopeBase, "exec", func(ctx *utils.ScopeContext) {

		exec_cmd := utils.GetMapValue(ctx, hc.getLxc(ctx).ContainerConfig, "exec").(string)
		call := make([]string, 0)
		if strings.Contains(exec_cmd, " ") {
			call = strings.Split(exec_cmd, " ")
		} else {
			call = []string{exec_cmd}
		}

		ctx.Must(hc.getLxc(ctx).Exec(ctx, call...))
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
