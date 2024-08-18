package hab

import (
	"github.com/tuxounet/k-hab/utils"
)

func (h *Hab) loadContainers(ctx *utils.ScopeContext) error {
	return ctx.Scope(h.scopeBase, "loadContainers", func(ctx *utils.ScopeContext) {
		for _, confContainer := range h.config.ContainersConfig {

			found := false
			for _, localContainer := range h.containers {
				if localContainer.name == confContainer.Name {
					found = true
					break
				}
			}
			if !found {
				container := newHabContainer(confContainer.Name, h)
				h.containers = append(h.containers, container)
			}
		}
	})
}

func (h *Hab) getEntryContainer(ctx *utils.ScopeContext) *HabContainer {
	return utils.ScopingWithReturn(ctx, h.scopeBase, "getEntryContainer", func(ctx *utils.ScopeContext) *HabContainer {

		entrypoint := utils.GetMapValue(ctx, h.config.HabConfig, "entry.container").(string)
		container := h.getContainer(ctx, entrypoint)
		if container == nil {
			ctx.Must(ctx.Error("Container not found"))
		}
		return container
	})
}

func (h *Hab) getContainer(ctx *utils.ScopeContext, name string) *HabContainer {
	return utils.ScopingWithReturn(ctx, h.scopeBase, "getContainer", func(ctx *utils.ScopeContext) *HabContainer {
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

func (h *Hab) upContainers(ctx *utils.ScopeContext) error {
	return ctx.Scope(h.scopeBase, "upContainers", func(ctx *utils.ScopeContext) {

		ctx.Must(h.loadContainers(ctx))
		for _, container := range h.containers {
			ctx.Must(container.provision(ctx))
		}
		ctx.Log.InfoF("Uped %d containers", len(h.containers))
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
	})
}

func (h *Hab) downContainers(ctx *utils.ScopeContext) error {
	return ctx.Scope(h.scopeBase, "downContainers", func(ctx *utils.ScopeContext) {
		ctx.Must(h.loadContainers(ctx))
		for _, container := range h.containers {
			ctx.Must(container.unprovision(ctx))
		}

	})
}
