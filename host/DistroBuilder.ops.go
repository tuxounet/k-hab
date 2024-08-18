package host

import (
	"os"
	"path"

	"github.com/tuxounet/k-hab/utils"
)

type DistroBuilderResult struct {
	Built           bool
	MetadataPackage string
	RootfsPackage   string
}

func (l *DistroBuilder) BuildDistro(ctx *utils.ScopeContext, name string, builderConfig string) *DistroBuilderResult {
	return utils.ScopingWithReturn(ctx, l.scopeBase, "BuildDistro", func(ctx *utils.ScopeContext) *DistroBuilderResult {

		distroFolder := path.Join(l.getImageBuildPath(ctx), name)
		os.MkdirAll(distroFolder, 0755)
		built := false
		if l.configHasChnaged(ctx, name, builderConfig) {
			ctx.Log.InfoF("Building distro %s", name)
			distroBuildFile := path.Join(distroFolder, "distro.yaml")
			ctx.Must(os.WriteFile(distroBuildFile, []byte(builderConfig), 0644))

			cmd := l.withDistroBuilderCmd(ctx, "build-lxd", distroBuildFile)
			cmd.Cwd = &distroFolder

			ctx.Must(utils.OsExec(ctx, cmd))

		}

		metadataPackage := path.Join(distroFolder, "incus.tar.xz")
		rootfsPackage := path.Join(distroFolder, "rootfs.squashfs")

		return &DistroBuilderResult{
			MetadataPackage: metadataPackage,
			RootfsPackage:   rootfsPackage,
			Built:           built,
		}
	})
}

func (l *DistroBuilder) RemoveCache(ctx *utils.ScopeContext, name string) error {
	return ctx.Scope(l.scopeBase, "RemoveCache", func(ctx *utils.ScopeContext) {
		distroFolder := path.Join(l.getImageBuildPath(ctx), name)
		cmd := utils.NewCmdCall("sudo", "rm", "-rf", distroFolder)
		ctx.Must(utils.OsExec(ctx, cmd))
	})
}

func (l *DistroBuilder) configHasChnaged(ctx *utils.ScopeContext, name string, expectedConfig string) bool {
	return utils.ScopingWithReturn(ctx, l.scopeBase, "configHasChnaged", func(ctx *utils.ScopeContext) bool {
		distroFolder := path.Join(l.getImageBuildPath(ctx), name)
		distroBuildFile := path.Join(distroFolder, "distro.yaml")

		if _, err := os.Stat(distroBuildFile); os.IsNotExist(err) {
			return true
		}
		metadataPackage := path.Join(distroFolder, "incus.tar.xz")
		rootfsPackage := path.Join(distroFolder, "rootfs.squashfs")

		if _, err := os.Stat(metadataPackage); os.IsNotExist(err) {
			return true
		}

		if _, err := os.Stat(rootfsPackage); os.IsNotExist(err) {
			return true
		}

		currentConfig, err := os.ReadFile(distroBuildFile)
		ctx.Must(err)

		return string(currentConfig) != expectedConfig
	})

}
