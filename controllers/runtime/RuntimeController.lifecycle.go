package runtime

import (
	"github.com/tuxounet/k-hab/controllers/dependencies"
	"github.com/tuxounet/k-hab/utils"
)

func (r *RuntimeController) IsPresent() (bool, error) {

	controller, err := r.ctx.GetController("DependenciesController")
	if err != nil {
		return false, err
	}
	dependencyController := controller.(*dependencies.DependenciesController)

	snapName := r.ctx.GetConfigValue("hab.lxd.snap")

	present, err := dependencyController.InstalledSnap(snapName)
	if err != nil {
		return false, err
	}
	return present, nil

}
func (r *RuntimeController) Provision() error {
	r.log.TraceF("Provisioning")

	controller, err := r.ctx.GetController("DependenciesController")
	if err != nil {
		return err
	}
	dependencyController := controller.(*dependencies.DependenciesController)

	snapName := r.ctx.GetConfigValue("hab.lxd.snap")
	snapMode := r.ctx.GetConfigValue("hab.lxd.snap_mode")
	present, err := dependencyController.InstalledSnap(snapName)
	if err != nil {
		return err
	}

	if !present {
		err := dependencyController.InstallSnap(snapName, snapMode)
		if err != nil {
			return err
		}
		r.log.DebugF("Provioned")
	}
	err = r.provisionStorage()
	if err != nil {
		return err
	}

	err = r.provisionNetwork()
	if err != nil {
		return err
	}

	err = r.provisionProfile()
	if err != nil {
		return err
	}

	r.log.DebugF("Provisioned")
	return nil
}

func (r *RuntimeController) Rm() error {

	present, err := r.IsPresent()
	if err != nil {
		return err
	}

	if present {
		cmd, err := r.withLxdCmd("shutdown")
		if err != nil {
			return err
		}
		err = utils.OsExec(cmd)
		if err != nil {
			return err
		}
		r.log.DebugF("Shutdowned")

	}
	return nil

}

func (r *RuntimeController) Unprovision() error {
	r.log.TraceF("Unprovisioning")

	present, err := r.IsPresent()
	if err != nil {
		return err
	}

	if present {

		err = r.unprovisionProfile()
		if err != nil {
			return err
		}

		err = r.unprovisionStorage()
		if err != nil {
			return err
		}

		controller, err := r.ctx.GetController("DependenciesController")
		if err != nil {
			return err
		}
		dependencyController := controller.(*dependencies.DependenciesController)
		snapName := r.ctx.GetConfigValue("hab.lxd.snap")
		err = dependencyController.RemoveSnap(snapName)
		if err != nil {
			return err
		}

	}

	r.log.DebugF("Unprovioned")
	return nil
}
func (r *RuntimeController) Nuke() error {
	r.log.TraceF("Nuking")
	err := r.nukeStorage()
	if err != nil {
		return err
	}

	controller, err := r.ctx.GetController("DependenciesController")
	if err != nil {
		return err
	}
	dependencyController := controller.(*dependencies.DependenciesController)

	snapName := r.ctx.GetConfigValue("hab.lxd.snap")

	err = dependencyController.RemoveSnapSnapshots(snapName)
	if err != nil {
		return err
	}

	r.log.DebugF("Nuked")
	return nil

}
