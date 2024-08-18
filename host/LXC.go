package host

import (
	"fmt"
	"runtime"
	"time"

	"github.com/tuxounet/k-hab/utils"
)

type LXC struct {
	scopeBase string
	name      string
	arch      string

	habConfig       map[string]interface{}
	ContainerConfig map[string]interface{}
}

func NewLXC(name string, habConfig map[string]interface{}, containerConfig map[string]interface{}) *LXC {

	return &LXC{
		scopeBase:       "LXC",
		name:            name,
		arch:            runtime.GOARCH,
		habConfig:       habConfig,
		ContainerConfig: containerConfig,
	}
}

func (l *LXC) withLxcCmd(ctx *utils.ScopeContext, args ...string) *utils.CmdCall {
	return utils.ScopingWithReturn(ctx, l.scopeBase, "Present", func(ctx *utils.ScopeContext) *utils.CmdCall {
		return utils.WithCmdCall(ctx, l.habConfig, "lxd.lxc.command.prefix", "lxd.lxc.command.name", args...)
	})
}

func (l *LXC) Present(ctx *utils.ScopeContext) bool {
	return utils.ScopingWithReturn(ctx, l.scopeBase, "Present", func(ctx *utils.ScopeContext) bool {
		out := utils.JsonCommandOutput[[]map[string]interface{}](ctx, l.withLxcCmd(ctx, "list", "--format", "json"))
		for _, container := range out {
			if container["name"] == l.name {
				return true
			}
		}
		return false
	})
}

func (l *LXC) Status(ctx *utils.ScopeContext) string {
	return utils.ScopingWithReturn(ctx, l.scopeBase, "Status", func(ctx *utils.ScopeContext) string {

		out := utils.JsonCommandOutput[[]map[string]interface{}](ctx, l.withLxcCmd(ctx, "list", "--format", "json"))

		for _, container := range out {
			if container["name"] == l.name {
				return container["status"].(string)
			}
		}
		return "unknown"
	})
}

func (l *LXC) Provision(ctx *utils.ScopeContext) error {
	return ctx.Scope(l.scopeBase, "Provision", func(ctx *utils.ScopeContext) {

		containerExists := l.Present(ctx)
		if !containerExists {

			containerImage := utils.GetMapValue(ctx, l.ContainerConfig, "image").(string)

			lxcProfile := utils.GetMapValue(ctx, l.habConfig, "lxd.lxc.profile").(string)
			lxdCmd := l.withLxcCmd(ctx, "init", containerImage, l.name, "--profile", lxcProfile)

			cloudInit := utils.GetMapValue(ctx, l.ContainerConfig, "cloud-init").(string)
			networkConfig := utils.GetMapValue(ctx, l.ContainerConfig, "network-config").(string)

			if cloudInit != "" {
				sCloudInit := utils.UnTemplate(ctx, cloudInit, map[string]interface{}{
					"hab":       l.habConfig,
					"container": l.ContainerConfig,
				})
				userDataInclude := fmt.Sprintf(`--config=user.user-data=%s`, sCloudInit)
				lxdCmd.Args = append(lxdCmd.Args, userDataInclude)
			}

			if networkConfig != "" {
				sNetworkConfig := utils.UnTemplate(ctx, networkConfig, map[string]interface{}{
					"hab":       l.habConfig,
					"container": l.ContainerConfig,
				})
				userDataInclude := fmt.Sprintf(`--config=user.network-config=%s`, sNetworkConfig)
				lxdCmd.Args = append(lxdCmd.Args, userDataInclude)
			}

			ctx.Must(utils.OsExec(ctx, lxdCmd))
		}
	})
}

func (l *LXC) Start(ctx *utils.ScopeContext) error {
	return ctx.Scope(l.scopeBase, "Start", func(ctx *utils.ScopeContext) {
		status := l.Status(ctx)

		if status != "Running" {
			ctx.Must(utils.OsExec(ctx, l.withLxcCmd(ctx, "start", l.name)))
		}

	})
}

func (l *LXC) WaitReady(ctx *utils.ScopeContext) error {
	return ctx.Scope(l.scopeBase, "WaitReady", func(ctx *utils.ScopeContext) {

		timeout := 30 * time.Second

		// Heure de fin
		heureFin := time.Now().Add(timeout)

		// Boucle jusqu'à l'heure de fin
		for time.Now().Before(heureFin) {

			status := l.Status(ctx)
			if status == "Running" {
				code := utils.OsExecWithExitCode(ctx, l.withLxcCmd(ctx, "exec", l.name, "--", "cloud-init", "status", "--wait"))
				if code == 2 || code == 0 {
					return
				}
			}
			time.Sleep(1 * time.Second)
		}
		ctx.Must(fmt.Errorf("Timeout to waiting ready"))
	})
}

func (l *LXC) Exec(ctx *utils.ScopeContext, command ...string) error {
	return ctx.Scope(l.scopeBase, "Exec", func(ctx *utils.ScopeContext) {
		cmd := l.withLxcCmd(ctx, "exec", l.name, "--")
		cmd.Args = append(cmd.Args, command...)
		ctx.Must(utils.OsExec(ctx, cmd))
	})
}

func (l *LXC) Stop(ctx *utils.ScopeContext) error {
	return ctx.Scope(l.scopeBase, "Stop", func(ctx *utils.ScopeContext) {
		status := l.Status(ctx)
		if status == "Running" {
			ctx.Must(utils.OsExec(ctx, l.withLxcCmd(ctx, "stop", l.name)))
		}
	})
}

func (l *LXC) Unprovision(ctx *utils.ScopeContext) error {
	return ctx.Scope(l.scopeBase, "Unprovision", func(ctx *utils.ScopeContext) {
		containerExists := l.Present(ctx)

		if containerExists {
			ctx.Must(l.Stop(ctx))
			ctx.Must(utils.OsExec(ctx, l.withLxcCmd(ctx, "delete", l.name)))
		}
	})
}
