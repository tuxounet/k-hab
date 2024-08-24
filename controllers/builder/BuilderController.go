package builder

import (
	"os"
	"path"
	"path/filepath"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/controllers/dependencies"
	"github.com/tuxounet/k-hab/utils"
)

type BuilderController struct {
	bases.BaseController
	ctx bases.IContext
}

func NewBuilderController(ctx bases.IContext) *BuilderController {
	return &BuilderController{
		ctx: ctx,
	}
}

func (b *BuilderController) Provision() error {
	controller, err := b.ctx.GetController("DependenciesController")
	if err != nil {
		return err
	}
	dependencyController := controller.(*dependencies.DependenciesController)

	config := b.ctx.GetHabConfig()

	snapName := utils.GetMapValue(config, "distrobuilder.snap").(string)
	snapMode := utils.GetMapValue(config, "distrobuilder.snap_mode").(string)
	present, err := dependencyController.InstalledSnap(snapName)
	if err != nil {
		return err
	}

	if !present {
		err := dependencyController.InstallSnap(snapName, snapMode)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *BuilderController) Unprovision() error {
	controller, err := b.ctx.GetController("DependenciesController")
	if err != nil {
		return err
	}
	dependencyController := controller.(*dependencies.DependenciesController)

	config := b.ctx.GetHabConfig()

	snapName := utils.GetMapValue(config, "distrobuilder.snap").(string)
	err = dependencyController.RemoveSnap(snapName)

	if err != nil {
		return err
	}

	err = dependencyController.RemoveSnapSnapshots(snapName)
	if err != nil {
		return err
	}
	return nil

}

func (b *BuilderController) Nuke() error {

	buildPath, err := b.getImageBuildPath()
	if err != nil {
		return err
	}
	err = utils.OsExec(utils.NewCmdCall("sudo", "rm", "-rf", buildPath))
	if err != nil {
		return err
	}

	return nil

}

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
