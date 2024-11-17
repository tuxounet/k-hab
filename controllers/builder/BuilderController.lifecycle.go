package builder

import (
	"github.com/tuxounet/k-hab/utils"
)

func (b *BuilderController) Install() error {
	b.log.TraceF("Installing")
	dependencyController, err := b.getDependenciesController()
	if err != nil {
		return err
	}

	snapName := b.ctx.GetConfigValue("hab.distrobuilder.snap")
	snapMode := b.ctx.GetConfigValue("hab.distrobuilder.snap_mode")
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
	b.log.DebugF("Installed")
	return nil
}

func (b *BuilderController) Uninstall() error {
	dependencyController, err := b.getDependenciesController()
	if err != nil {
		return err
	}

	snapName := b.ctx.GetConfigValue("hab.distrobuilder.snap")
	err = dependencyController.RemoveSnap(snapName)

	if err != nil {
		return err
	}

	return nil
}

func (b *BuilderController) Nuke() error {
	dependencyController, err := b.getDependenciesController()
	if err != nil {
		return err
	}

	snapName := b.ctx.GetConfigValue("hab.distrobuilder.snap")

	err = dependencyController.RemoveSnapSnapshots(snapName)
	if err != nil {
		return err
	}

	buildPath, err := b.getImageBuildPath()
	if err != nil {
		return err
	}

	cmd, err := utils.WithCmdCall(b.ctx, "hab.commands.rm.prefix", "hab.commands.rm", "-rf", buildPath)
	if err != nil {
		return err
	}
	err = utils.OsExec(cmd)
	if err != nil {
		return err
	}

	return nil

}
