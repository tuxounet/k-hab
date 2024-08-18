package host

import (
	"os"
	"path"
	"path/filepath"

	"github.com/tuxounet/k-hab/utils"
)

func (l *LXD) getStoragePath(ctx *utils.ScopeContext) string {

	return utils.ScopingWithReturn(ctx, l.scopeBase, "getStoragePath", func(ctx *utils.ScopeContext) string {

		storagePathDefinition := utils.GetMapValue(ctx, l.habConfig, "lxd.lxc.storage.path").(string)
		var storagePath string
		isAbsolute := filepath.IsAbs(storagePathDefinition)
		if !isAbsolute {
			cwd, err := os.Getwd()
			ctx.Must(err)
			storagePath = path.Join(cwd, storagePathDefinition)

		} else {
			storagePath = storagePathDefinition
		}
		os.MkdirAll(storagePath, 0755)
		return storagePath

	})
}

func (l *LXD) ProvisionStorage(ctx *utils.ScopeContext) error {
	return ctx.Scope(l.scopeBase, "ProvisionStorage", func(ctx *utils.ScopeContext) {

		storage_pool := utils.GetMapValue(ctx, l.habConfig, "lxd.lxc.storage.pool").(string)
		storage_driver := utils.GetMapValue(ctx, l.habConfig, "lxd.lxc.storage.driver").(string)
		storage_path := l.getStoragePath(ctx)

		storageExists := l.existsStorage(ctx, storage_pool)

		if !storageExists {
			ctx.Must(l.createStorage(ctx, storage_pool, storage_driver, "source="+storage_path))
		}

	})
}

func (l *LXD) UnprovisionStorage(ctx *utils.ScopeContext) error {
	return ctx.Scope(l.scopeBase, "UnprovisionStorage", func(ctx *utils.ScopeContext) {
		storage_pool := utils.GetMapValue(ctx, l.habConfig, "lxd.lxc.storage.pool").(string)

		storageExists := l.existsStorage(ctx, storage_pool)
		if storageExists {
			ctx.Must(l.removeStorage(ctx, storage_pool))
		}

	})
}

func (l *LXD) NukeStorage(ctx *utils.ScopeContext) error {
	return ctx.Scope(l.scopeBase, "NukeStorage", func(ctx *utils.ScopeContext) {
		storage_path := l.getStoragePath(ctx)
		ctx.Must(utils.OsExec(ctx, utils.NewCmdCall("sudo", "rm", "-rvf", storage_path)))
	})
}

func (l *LXD) existsStorage(ctx *utils.ScopeContext, name string) bool {

	return utils.ScopingWithReturn(ctx, l.scopeBase, "existsStorage", func(ctx *utils.ScopeContext) bool {

		arr := utils.JsonCommandOutput[[]map[string]interface{}](ctx, l.withLxcCmd(ctx, "storage", "ls", "--format", "json"))

		for _, profile := range arr {
			if profile["name"] == name {
				return true
			}
		}

		return false
	})
}

func (l *LXD) createStorage(ctx *utils.ScopeContext, name string, driver string, options ...string) error {
	return ctx.Scope(l.scopeBase, "createStorage", func(ctx *utils.ScopeContext) {
		ctx.Must(l.NukeStorage(ctx))
		ctx.Must(os.MkdirAll(l.getStoragePath(ctx), 0755))

		cmd := l.withLxcCmd(ctx, "storage", "create", name, driver)
		cmd.Args = append(cmd.Args, options...)

		ctx.Must(utils.OsExec(ctx, cmd))
	})
}

func (l *LXD) removeStorage(ctx *utils.ScopeContext, name string) error {
	return ctx.Scope(l.scopeBase, "removeStorage", func(ctx *utils.ScopeContext) {

		ctx.Must(utils.OsExec(ctx, l.withLxcCmd(ctx, "storage", "delete", name)))
	})
}
