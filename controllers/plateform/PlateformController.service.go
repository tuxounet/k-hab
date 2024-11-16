package plateform

func (r *PlateformController) presentService() (bool, error) {

	dependencyController, err := r.getDependenciesController()
	if err != nil {
		return false, err
	}

	snapName := r.ctx.GetConfigValue("hab.plateform.snap")
	present, err := dependencyController.InstalledSnap(snapName)
	if err != nil {
		return false, err
	}
	return present, nil
}

func (r *PlateformController) provisionService() error {
	dependencyController, err := r.getDependenciesController()
	if err != nil {
		return err
	}

	snapName := r.ctx.GetConfigValue("hab.plateform.snap")
	snapMode := r.ctx.GetConfigValue("hab.plateform.snap_mode")

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
	r.log.DebugF("Provisioned")
	return nil
}

func (r *PlateformController) unprovisionService() error {

	dependencyController, err := r.getDependenciesController()
	if err != nil {
		return err
	}

	snapName := r.ctx.GetConfigValue("hab.plateform.snap")
	err = dependencyController.RemoveSnap(snapName)

	if err != nil {
		return err
	}

	return nil

}

func (r *PlateformController) nukeService() error {
	dependencyController, err := r.getDependenciesController()
	if err != nil {
		return err
	}

	snapName := r.ctx.GetConfigValue("hab.plateform.snap")

	err = dependencyController.RemoveSnapSnapshots(snapName)
	if err != nil {
		return err
	}

	return nil

}
