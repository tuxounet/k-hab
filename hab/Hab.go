package hab

import (
	"github.com/tuxounet/k-hab/config"
	"github.com/tuxounet/k-hab/host"
	"github.com/tuxounet/k-hab/utils"
)

type Hab struct {
	scopeBase   string
	ctx         *utils.ScopeContext
	config      *config.Config
	lxd         *host.LXD
	builder     *host.DistroBuilder
	containers  []*HabContainer
	images      []*HabImage
	httpEgress  *host.HttpEgress
	httpIngress *host.HttpIngress
}

func NewHab(quiet bool) *Hab {
	scopeCtx := utils.NewScopeContext(quiet, "Hab")
	config := config.NewConfig()
	config.Load(scopeCtx)

	hab := &Hab{
		scopeBase:   scopeCtx.Name,
		ctx:         scopeCtx,
		config:      config,
		builder:     host.NewDistroBuilder(config.HabConfig),
		lxd:         host.NewLXD(config.HabConfig),
		httpEgress:  host.NewHttpEgress(config.HabConfig),
		httpIngress: host.NewHttpIngress(config.HabConfig),
		containers:  make([]*HabContainer, 0),
	}

	return hab
}

func (h *Hab) Provision() error {
	return h.ctx.Scope(h.scopeBase, "Provision", func(ctx *utils.ScopeContext) {

		ctx.Must(h.provisionImages(ctx))
		ctx.Must(h.provisionContainers(ctx))

	})

}

func (h *Hab) Start() error {
	return h.ctx.Scope(h.scopeBase, "Start", func(ctx *utils.ScopeContext) {

		//Ensure Provisioning
		ctx.Must(h.provisionImages(ctx))
		ctx.Must(h.provisionContainers(ctx))

		//Start
		ctx.Must(h.httpEgress.Start(ctx))
		ctx.Must(h.httpIngress.Start(ctx))
		ctx.Must(h.startContainers(ctx))

		//launch
		entrypoint := utils.GetMapValue(ctx, h.config.HabConfig, "entry.container").(string)
		container := h.getContainer(ctx, entrypoint)
		if container == nil {
			ctx.Must(ctx.Error("Container not found"))
		}
		ctx.Must(container.waitReady(ctx))
		ctx.Must(container.exec(ctx))

		//cleanup
		ctx.Must(h.httpIngress.Stop(ctx))
		ctx.Must(h.httpEgress.Stop(ctx))
		ctx.Must(h.stopContainers(ctx))
	})

}
func (h *Hab) Shell() error {
	return h.ctx.Scope(h.scopeBase, "Shell", func(ctx *utils.ScopeContext) {

		//Ensure Provisioning
		ctx.Must(h.provisionImages(ctx))
		ctx.Must(h.provisionContainers(ctx))

		//Start
		ctx.Must(h.httpEgress.Start(ctx))
		ctx.Must(h.httpIngress.Start(ctx))
		ctx.Must(h.startContainers(ctx))

		//launch
		entrypoint := utils.GetMapValue(ctx, h.config.HabConfig, "entry.container").(string)
		container := h.getContainer(ctx, entrypoint)
		if container == nil {
			ctx.Must(ctx.Error("Container not found"))
		}

		ctx.Must(container.shell(ctx))

		//cleanup
		ctx.Must(h.httpIngress.Stop(ctx))
		ctx.Must(h.httpEgress.Stop(ctx))
		ctx.Must(h.stopContainers(ctx))
	})

}
func (h *Hab) Stop() error {
	return h.ctx.Scope(h.scopeBase, "Stop", func(ctx *utils.ScopeContext) {
		ctx.Must(h.stopContainers(ctx))
		ctx.Must(h.lxd.Down(ctx))
	})
}

func (h *Hab) Rm() error {
	return h.ctx.Scope(h.scopeBase, "Rm", func(ctx *utils.ScopeContext) {
		ctx.Must(h.stopContainers(ctx))
		ctx.Must(h.unprovisionContainers(ctx))
		ctx.Must(h.unprovisionImages(ctx))
		ctx.Must(h.lxd.Down(ctx))
	})
}

func (h *Hab) Unprovision() error {
	return h.ctx.Scope(h.scopeBase, "Unprovision", func(ctx *utils.ScopeContext) {
		lxdPresent := h.lxd.Present(ctx)
		if lxdPresent {
			ctx.Must(h.stopContainers(ctx))
			ctx.Must(h.unprovisionContainers(ctx))
			ctx.Must(h.lxd.Unprovision(ctx))
			ctx.Must(h.lxd.Down(ctx))
			ctx.Must(h.unprovisionImages(ctx))
		}

	})
}

func (h *Hab) Nuke() error {
	return h.ctx.Scope(h.scopeBase, "Nuke", func(ctx *utils.ScopeContext) {
		ctx.Must(h.lxd.Nuke(ctx))
		ctx.Must(h.builder.Nuke(ctx))
	})
}
