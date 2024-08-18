package hab

import (
	"github.com/tuxounet/k-hab/config"
	"github.com/tuxounet/k-hab/utils"
)

type HabImage struct {
	scopeBase string
	name      string
	hab       *Hab
	config    *config.HabImageConfig
}

func newHabImage(name string, hab *Hab, config *config.HabImageConfig) *HabImage {
	base := "HabImage/" + name

	return &HabImage{
		scopeBase: base,
		name:      name,
		hab:       hab,
		config:    config,
	}
}

func (hi *HabImage) present(ctx *utils.ScopeContext) bool {
	return utils.ScopingWithReturnOnly(ctx, hi.scopeBase, "present", func(ctx *utils.ScopeContext) bool {
		return hi.hab.lxd.PresentImage(ctx, hi.name)

	})
}

func (hi *HabImage) provision(ctx *utils.ScopeContext) error {
	return ctx.Scope(hi.scopeBase, "provision", func(ctx *utils.ScopeContext) {

		buildResult := hi.hab.builder.BuildDistro(ctx, hi.name, hi.config.Builder)

		ctx.Must(hi.hab.lxd.RegisterImage(ctx, hi.name, buildResult.MetadataPackage, buildResult.RootfsPackage, buildResult.Built))

	})
}

func (hi *HabImage) unprovision(ctx *utils.ScopeContext) error {
	return ctx.Scope(hi.scopeBase, "unprovision", func(ctx *utils.ScopeContext) {
		ctx.Must(hi.hab.lxd.RemoveImage(ctx, hi.name))

	})
}

func (hi *HabImage) nuke(ctx *utils.ScopeContext, name string) error {
	return ctx.Scope(hi.scopeBase, "nuke", func(ctx *utils.ScopeContext) {

		ctx.Must(hi.hab.builder.RemoveCache(ctx, hi.name))
	})
}
