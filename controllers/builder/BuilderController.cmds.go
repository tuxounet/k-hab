package builder

import (
	"os"
	"path"
	"path/filepath"

	"github.com/tuxounet/k-hab/utils"
)

func (b *BuilderController) getImageBuildPath() (string, error) {
	config := b.ctx.GetHabConfig()
	buildPathDefinition := utils.GetMapValue(config, "distrobuilder.build.path").(string)
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
	habConfig := b.ctx.GetHabConfig()
	return utils.WithCmdCall(habConfig, "distrobuilder.command.prefix", "distrobuilder.command.name", args...)
}
