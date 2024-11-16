package builder

import (
	"os"
	"path"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/controllers/dependencies"
	"github.com/tuxounet/k-hab/utils"
)

func (b *BuilderController) getImageBuildPath() (string, error) {

	buildPathDefinition := b.ctx.GetConfigValue("hab.distrobuilder.build.path")

	storageRoot, err := b.ctx.GetStorageRoot()
	if err != nil {
		return "", err
	}
	buildPath := path.Join(storageRoot, buildPathDefinition)

	os.MkdirAll(buildPath, 0755)
	return buildPath, nil

}
func (b *BuilderController) withDistroBuilderCmd(args ...string) (*utils.CmdCall, error) {
	return utils.WithCmdCall(b.ctx, "hab.distrobuilder.command.prefix", "hab.distrobuilder.command.name", args...)
}

func (l *BuilderController) getDependenciesController() (*dependencies.DependenciesController, error) {
	controller, err := l.ctx.GetController(bases.DependenciesController)
	if err != nil {
		return nil, err
	}
	dependencyController := controller.(*dependencies.DependenciesController)
	return dependencyController, nil
}
