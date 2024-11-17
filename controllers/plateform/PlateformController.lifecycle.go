package plateform

import (
	"github.com/tuxounet/k-hab/utils"
)

func (r *PlateformController) Install() error {
	r.log.TraceF("Installing")

	err := r.installService()
	if err != nil {
		return err
	}

	r.log.DebugF("Installed")
	return nil
}

func (r *PlateformController) Uninstall() error {
	r.log.TraceF("Uninstalling")

	err := r.uninstallService()
	if err != nil {
		return err
	}

	r.log.DebugF("Uninstalled")
	return nil
}

func (r *PlateformController) Provision() error {
	r.log.TraceF("Provisioning")

	err := r.provisionStorage()
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

	present, err := r.presentService()
	if err != nil {
		return err
	}

	if present {
		cmd, err := r.withLxcCmd("stop", "--all")
		if err != nil {
			return err
		}
		err = utils.OsExec(cmd)
		if err != nil {
			return err
		}
		r.log.DebugF("Stopped")

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
	return nil

}

func (r *PlateformController) Unprovision() error {
	r.log.TraceF("Unprovisioning")

	err := r.Rm()
	if err != nil {
		return err
	}

	r.log.DebugF("Unprovioned")
	return nil
}
func (r *PlateformController) Nuke() error {
	r.log.TraceF("Nuking")
	err := r.uninstallService()
	if err != nil {
		return err
	}

	err = r.nukeService()
	if err != nil {
		return err
	}

	err = r.nukeStorage()
	if err != nil {
		return err
	}

	r.log.DebugF("Nuked")
	return nil

}
