package builder

import (
	"github.com/tuxounet/k-hab/controllers/dependencies"
	"github.com/tuxounet/k-hab/utils"
)

func (b *BuilderController) Provision() error {
	b.log.TraceF("Provisioning")
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
	b.log.DebugF("Provisioned")
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

	return nil

}

func (b *BuilderController) Nuke() error {
	controller, err := b.ctx.GetController("DependenciesController")
	if err != nil {
		return err
	}
	dependencyController := controller.(*dependencies.DependenciesController)

	config := b.ctx.GetHabConfig()

	snapName := utils.GetMapValue(config, "distrobuilder.snap").(string)

	err = dependencyController.RemoveSnapSnapshots(snapName)
	if err != nil {
		return err
	}

	buildPath, err := b.getImageBuildPath()
	if err != nil {
		return err
	}

	cmd, err := utils.WithCmdCall(b.ctx.GetHabConfig(), "rm.prefix", "rm.name", "-rf", buildPath)
	if err != nil {
		return err
	}
	err = utils.OsExec(cmd)
	if err != nil {
		return err
	}

	return nil

}
