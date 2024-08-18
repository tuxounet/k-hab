package hab

import (
	"github.com/tuxounet/k-hab/utils"
)

func (h *Hab) loadContainers(ctx *utils.ScopeContext) error {
	return ctx.Scope(h.scopeBase, "loadContainers", func(ctx *utils.ScopeContext) {
		for _, confContainer := range h.config.ContainersConfig {

			name := confContainer.(map[string]interface{})["name"].(string)
			found := false
			for _, localContainer := range h.containers {
				if localContainer.name == name {
					found = true
					break
				}
			}
			if !found {
				container := newHabContainer(name, h)
				h.containers = append(h.containers, container)
			}
		}
	})
}
func (h *Hab) getContainer(ctx *utils.ScopeContext, name string) *HabContainer {
	return utils.ScopingWithReturnOnly(ctx, h.scopeBase, "getContainer", func(ctx *utils.ScopeContext) *HabContainer {
		ctx.Must(h.loadContainers(ctx))
		for _, container := range h.containers {
			if container.name == name {
				return container
			}
		}
		ctx.Must(ctx.Error("Container not found"))
		return nil
	})
}

func (h *Hab) provisionContainers(ctx *utils.ScopeContext) error {
	return ctx.Scope(h.scopeBase, "provisionContainers", func(ctx *utils.ScopeContext) {
		ctx.Must(h.lxd.Provision(ctx))
		ctx.Must(h.loadContainers(ctx))
		for _, container := range h.containers {
			ctx.Must(container.provision(ctx))
		}
		ctx.Log.InfoF("Provisioned %d containers", len(h.containers))
	})
}

func (h *Hab) startContainers(ctx *utils.ScopeContext) error {
	return ctx.Scope(h.scopeBase, "startContainers", func(ctx *utils.ScopeContext) {
		ctx.Must(h.loadContainers(ctx))
		ctx.Must(h.lxd.Up(ctx))
		for _, container := range h.containers {
			ctx.Must(container.up(ctx))
		}
	})
}

func (h *Hab) stopContainers(ctx *utils.ScopeContext) error {
	return ctx.Scope(h.scopeBase, "stopContainers", func(ctx *utils.ScopeContext) {
		ctx.Must(h.loadContainers(ctx))
		for _, container := range h.containers {
			ctx.Must(container.down(ctx))
		}
		ctx.Must(h.lxd.Down(ctx))
	})
}

func (h *Hab) unprovisionContainers(ctx *utils.ScopeContext) error {
	return ctx.Scope(h.scopeBase, "unprovisionContainers", func(ctx *utils.ScopeContext) {
		ctx.Must(h.loadContainers(ctx))
		for _, container := range h.containers {
			ctx.Must(container.unprovision(ctx))
		}

	})
}
