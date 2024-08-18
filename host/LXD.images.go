package host

import (
	"github.com/tuxounet/k-hab/utils"
)

func (l *LXD) PresentImage(ctx *utils.ScopeContext, name string) bool {
	return utils.ScopingWithReturn(ctx, l.scopeBase, "PresentImage", func(ctx *utils.ScopeContext) bool {

		out := utils.CommandSyncJsonArrayOutput(ctx, l.withLxcCmd(ctx, "image", "list", "--format", "json"))

		for _, image := range out {
			aliases := image["aliases"].([]interface{})
			for _, ali := range aliases {
				alias := ali.(map[string]interface{})
				if alias["name"].(string) == name {
					return true
				}
			}
		}
		return false
	})
}

func (l *LXD) RegisterImage(ctx *utils.ScopeContext, name string, metadataPackage string, rootfsPackage string, force bool) error {
	return ctx.Scope(l.scopeBase, "RegisterImage", func(ctx *utils.ScopeContext) {

		present := l.PresentImage(ctx, name)
		if present || force {
			ctx.Must(l.RemoveImage(ctx, name))
		}

		ctx.Must(utils.ExecSyncOutput(ctx, l.withLxcCmd(ctx, "image", "import", metadataPackage, rootfsPackage, "--alias", name)))
	})
}

func (l *LXD) RemoveImage(ctx *utils.ScopeContext, name string) error {
	return ctx.Scope(l.scopeBase, "RemoveImage", func(ctx *utils.ScopeContext) {

		present := l.PresentImage(ctx, name)
		if present {
			ctx.Must(utils.ExecSyncOutput(ctx, l.withLxcCmd(ctx, "image", "delete", name)))
		}
	})
}
