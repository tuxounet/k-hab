package host

import (
	"os"
	"path"
	"path/filepath"

	"github.com/tuxounet/k-hab/config"
	"github.com/tuxounet/k-hab/utils"
)

type DistroBuilder struct {
	scopeBase string
	cwd       string
	habConfig config.HabConfig
}

func NewDistroBuilder(habConfig config.HabConfig, cwd string) *DistroBuilder {

	return &DistroBuilder{
		scopeBase: "DistroBuilder",
		cwd:       cwd,
		habConfig: habConfig,
	}
}
func (l *DistroBuilder) withDistroBuilderCmd(ctx *utils.ScopeContext, args ...string) *utils.CmdCall {
	return utils.ScopingWithReturn(ctx, l.scopeBase, "withDistroBuilderCmd", func(ctx *utils.ScopeContext) *utils.CmdCall {
		return utils.WithCmdCall(ctx, l.habConfig, "distrobuilder.command.prefix", "distrobuilder.command.name", args...)
	})
}

func (l *DistroBuilder) getImageBuildPath(ctx *utils.ScopeContext) string {

	return utils.ScopingWithReturn(ctx, l.scopeBase, "getImageBuildPath", func(ctx *utils.ScopeContext) string {

		buildPathDefinition := utils.GetMapValue(ctx, l.habConfig, "distrobuilder.build.path").(string)
		var buildPath string
		isAbsolute := filepath.IsAbs(buildPathDefinition)
		if !isAbsolute {
			cwd, err := os.Getwd()
			ctx.Must(err)
			buildPath = path.Join(cwd, buildPathDefinition)

		} else {
			buildPath = buildPathDefinition
		}
		os.MkdirAll(buildPath, 0755)
		return buildPath

	})
}

func (l *DistroBuilder) Present(ctx *utils.ScopeContext) bool {
	return utils.ScopingWithReturn(ctx, l.scopeBase, "Present", func(ctx *utils.ScopeContext) bool {
		snaps := NewSnapPackages(l.habConfig)

		snapName := utils.GetMapValue(ctx, l.habConfig, "distrobuilder.snap").(string)
		present := snaps.InstalledSnap(ctx, snapName)

		return present
	})
}

func (l *DistroBuilder) Provision(ctx *utils.ScopeContext) error {

	return ctx.Scope(l.scopeBase, "Provision", func(ctx *utils.ScopeContext) {

		snaps := NewSnapPackages(l.habConfig)
		snapName := utils.GetMapValue(ctx, l.habConfig, "distrobuilder.snap").(string)
		snapMode := utils.GetMapValue(ctx, l.habConfig, "distrobuilder.snap_mode").(string)
		present := snaps.InstalledSnap(ctx, snapName)

		if !present {
			ctx.Must(snaps.InstallSnap(ctx, snapName, snapMode))
		}

	})
}
func (l *DistroBuilder) Unprovision(ctx *utils.ScopeContext) error {
	return ctx.Scope(l.scopeBase, "Build", func(ctx *utils.ScopeContext) {

		snaps := NewSnapPackages(l.habConfig)
		snapName := utils.GetMapValue(ctx, l.habConfig, "distrobuilder.snap").(string)
		ctx.Must(snaps.RemoveSnap(ctx, snapName))
		ctx.Must(snaps.RemoveSnapSnapshots(ctx, snapName))
	})
}

func (l *DistroBuilder) Nuke(ctx *utils.ScopeContext) error {
	return ctx.Scope(l.scopeBase, "Nuke", func(ctx *utils.ScopeContext) {

		buildPath := l.getImageBuildPath(ctx)
		ctx.Must(utils.OsExec(ctx, utils.NewCmdCall("sudo", "rm", "-rf", buildPath)))
	})
}
