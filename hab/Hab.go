package hab

import (
	"os"

	"github.com/tuxounet/k-hab/config"
	"github.com/tuxounet/k-hab/host"
	"github.com/tuxounet/k-hab/utils"
)

type Hab struct {
	scopeBase   string
	cwd         string
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

	cwd, err := os.Getwd()
	scopeCtx.Must(err)

	config := config.NewConfig()
	config.Load(scopeCtx)

	hab := &Hab{
		scopeBase:   scopeCtx.Name,
		cwd:         cwd,
		ctx:         scopeCtx,
		config:      config,
		builder:     host.NewDistroBuilder(config.HabConfig, cwd),
		lxd:         host.NewLXD(config.HabConfig, cwd),
		httpEgress:  host.NewHttpEgress(config.HabConfig),
		httpIngress: host.NewHttpIngress(config.HabConfig),
		containers:  make([]*HabContainer, 0),
	}

	return hab
}

func (h *Hab) Provision() error {
	return h.ctx.Scope(h.scopeBase, "Provision", func(ctx *utils.ScopeContext) {
		ctx.Must(h.builder.Provision(ctx))
		ctx.Must(h.lxd.Provision(ctx))

	})

}

func (h *Hab) Start() error {
	return h.ctx.Scope(h.scopeBase, "Start", func(ctx *utils.ScopeContext) {

		//Ensure Provisioning
		ctx.Must(h.Provision())

		//Start
		ctx.Must(h.httpEgress.Start(ctx))

		ctx.Must(h.upImages(ctx))
		ctx.Must(h.upContainers(ctx))

		ctx.Must(h.httpIngress.Start(ctx))
		ctx.Must(h.startContainers(ctx))
	})
}

func (h *Hab) Up() error {

	return h.ctx.Scope(h.scopeBase, "Up", func(ctx *utils.ScopeContext) {
		ctx.Must(h.Start())

		container := h.getEntryContainer(ctx)
		ctx.Must(container.waitReady(ctx))
		ctx.Must(container.entry(ctx))

		ctx.Must(h.Stop())
	})
}

func (h *Hab) Shell() error {
	return h.ctx.Scope(h.scopeBase, "Shell", func(ctx *utils.ScopeContext) {

		ctx.Must(h.Start())

		container := h.getEntryContainer(ctx)
		ctx.Must(container.shell(ctx))

		//cleanup
		ctx.Must(h.Stop())
	})

}

func (h *Hab) Stop() error {
	return h.ctx.Scope(h.scopeBase, "Stop", func(ctx *utils.ScopeContext) {
		ctx.Must(h.httpIngress.Stop(ctx))

		ctx.Must(h.stopContainers(ctx))

		ctx.Must(h.httpEgress.Stop(ctx))
	})
}

func (h *Hab) Rm() error {
	return h.ctx.Scope(h.scopeBase, "Rm", func(ctx *utils.ScopeContext) {
		ctx.Must(h.Stop())
		ctx.Must(h.downContainers(ctx))
		ctx.Must(h.downImages(ctx))
		ctx.Must(h.lxd.Down(ctx))
	})
}

func (h *Hab) Unprovision() error {
	return h.ctx.Scope(h.scopeBase, "Unprovision", func(ctx *utils.ScopeContext) {
		lxdPresent := h.lxd.Present(ctx)
		if lxdPresent {
			ctx.Must(h.Rm())
			ctx.Must(h.lxd.Unprovision(ctx))
		}

		builderPresent := h.builder.Present(ctx)
		if builderPresent {
			ctx.Must(h.builder.Unprovision(ctx))
		}

	})
}

func (h *Hab) Nuke() error {
	return h.ctx.Scope(h.scopeBase, "Nuke", func(ctx *utils.ScopeContext) {
		ctx.Must(h.Unprovision())
		ctx.Must(h.nukeImages(ctx))
		ctx.Must(h.lxd.Nuke(ctx))
		ctx.Must(h.builder.Nuke(ctx))
	})
}
