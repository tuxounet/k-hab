package plateform

import (
	"github.com/tuxounet/k-hab/controllers/dependencies"
	"github.com/tuxounet/k-hab/utils"
)

func (r *PlateformController) IsServerPresent() (bool, error) {

	controller, err := r.ctx.GetController("DependenciesController")
	if err != nil {
		return false, err
	}
	dependencyController := controller.(*dependencies.DependenciesController)

	aptName := r.ctx.GetConfigValue("hab.incus.apt.server")

	present, err := dependencyController.InstalledAPT(aptName)
	if err != nil {
		return false, err
	}
	return present, nil

}
func (r *PlateformController) IsClientPresent() (bool, error) {

	controller, err := r.ctx.GetController("DependenciesController")
	if err != nil {
		return false, err
	}
	dependencyController := controller.(*dependencies.DependenciesController)

	aptName := r.ctx.GetConfigValue("hab.incus.apt.client")

	present, err := dependencyController.InstalledAPT(aptName)
	if err != nil {
		return false, err
	}
	return present, nil

}
func (r *PlateformController) Provision() error {
	r.log.TraceF("Provisioning")

	err := r.provisionServer()
	if err != nil {
		return err
	}

	err = r.provisionClient()
	if err != nil {
		return err
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

func (r *PlateformController) Rm() error {

	present, err := r.IsClientPresent()
	if err != nil {
		return err
	}

	if present {
		cmd, err := r.withIncusCmd("stop", "--all")
		if err != nil {
			return err
		}
		err = utils.OsExec(cmd)
		if err != nil {
			return err
		}
		r.log.DebugF("Stopped")

	}
	return nil

}

func (r *PlateformController) Unprovision() error {
	r.log.TraceF("Unprovisioning")

	present, err := r.IsClientPresent()
	if err != nil {
		return err
	}

	if present {
		err = r.unprovisionProfile()
		if err != nil {
			return err
		}
		err = r.unprovisionNetwork()
		if err != nil {
			return err
		}

		err = r.unprovisionStorage()
		if err != nil {
			return err
		}
	}

	err = r.unprovisionClient()
	if err != nil {
		return err
	}

	err = r.unprovisionServer()
	if err != nil {
		return err
	}

	r.log.DebugF("Unprovioned")
	return nil
}
func (r *PlateformController) Nuke() error {
	r.log.TraceF("Nuking")
	err := r.nukeStorage()
	if err != nil {
		return err
	}

	r.log.DebugF("Nuked")
	return nil

}
