package plateform

import (
	"strings"

	"github.com/tuxounet/k-hab/utils"
)

func (r *PlateformController) provisionServer() error {
	dependencyController, err := r.getDependenciesController()
	if err != nil {
		return err
	}

	aptName := r.ctx.GetConfigValue("hab.incus.apt.server")

	present, err := dependencyController.InstalledAPT(aptName)
	if err != nil {
		return err
	}

	if !present {
		err := dependencyController.InstallAPT(aptName)
		if err != nil {
			return err
		}
		r.log.DebugF("%s provisionned", aptName)
	}
	return r.provisionDnsMasq()
}

func (r *PlateformController) provisionDnsMasq() error {
	dependencyController, err := r.getDependenciesController()
	if err != nil {
		return err
	}

	aptName := r.ctx.GetConfigValue("hab.incus.apt.dnsmasq")

	present, err := dependencyController.InstalledAPT(aptName)
	if err != nil {
		return err
	}

	if !present {
		err := dependencyController.InstallAPT(aptName)
		if err != nil {
			return err
		}
		r.log.DebugF("%s provisionned", aptName)
	}
	return nil
}

func (r *PlateformController) provisionClient() error {
	dependencyController, err := r.getDependenciesController()
	if err != nil {
		return err
	}

	aptName := r.ctx.GetConfigValue("hab.incus.apt.client")

	present, err := dependencyController.InstalledAPT(aptName)
	if err != nil {
		return err
	}

	if !present {
		err := dependencyController.InstallAPT(aptName)
		if err != nil {
			return err
		}
		r.log.DebugF("%s provisionned", aptName)
	}
	return nil
}

func (r *PlateformController) unprovisionServer() error {

	cmd, err := r.withPsCmd("aux")
	if err != nil {
		return err
	}
	out, err := utils.RawCommandOutput(cmd)
	if err != nil {
		return err
	}

	active := strings.Contains(out, "/incusd")
	if active {
		cmd, err := r.withSystemCtlCmd("stop", "incus")
		if err != nil {
			return err
		}
		err = utils.OsExec(cmd)
		if err != nil {
			return nil
		}

	}

	dependencyController, err := r.getDependenciesController()
	if err != nil {
		return err
	}
	aptName := r.ctx.GetConfigValue("hab.incus.apt.server")

	present, err := dependencyController.InstalledAPT(aptName)
	if err != nil {
		return err
	}
	if present {
		err := dependencyController.RemoveAPT(aptName)
		if err != nil {
			return err
		}

		r.log.DebugF("unprovisioned %s", aptName)
	}

	return r.unprovisionDnsmasq()

}
func (r *PlateformController) unprovisionDnsmasq() error {

	dependencyController, err := r.getDependenciesController()
	if err != nil {
		return err
	}
	aptName := r.ctx.GetConfigValue("hab.incus.apt.dnsmasq")

	present, err := dependencyController.InstalledAPT(aptName)
	if err != nil {
		return err
	}
	if present {
		err := dependencyController.RemoveAPT(aptName)
		if err != nil {
			return err
		}

		r.log.DebugF("unprovisioned %s", aptName)
	}

	return nil

}
func (r *PlateformController) unprovisionClient() error {
	dependencyController, err := r.getDependenciesController()
	if err != nil {
		return err
	}

	aptName := r.ctx.GetConfigValue("hab.incus.apt.client")

	present, err := dependencyController.InstalledAPT(aptName)
	if err != nil {
		return err
	}
	if present {
		err := dependencyController.RemoveAPT(aptName)
		if err != nil {
			return err
		}
		r.log.DebugF("Unprovioned %s", aptName)
	}
	return nil

}
