package builder

import (
	"os"
	"path"
	"path/filepath"

	"github.com/tuxounet/k-hab/utils"
)

func (b *BuilderController) getImageBuildPath() (string, error) {

	buildPathDefinition := b.ctx.GetConfigValue("hab.distrobuilder.build.path")
	var buildPath string
	isAbsolute := filepath.IsAbs(buildPathDefinition)
	if !isAbsolute {
		cwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		buildPath = path.Join(cwd, buildPathDefinition)

	} else {
		buildPath = buildPathDefinition
	}
	os.MkdirAll(buildPath, 0755)
	return buildPath, nil

}
func (b *BuilderController) withDistroBuilderCmd(args ...string) (*utils.CmdCall, error) {
	return utils.WithCmdCall(b.ctx, "hab.distrobuilder.command.prefix", "hab.distrobuilder.command.name", args...)
}
