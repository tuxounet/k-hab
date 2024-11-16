package containers

func (c *ContainersController) Provision() error {

	c.log.TraceF("Provisioning containers")
	err := c.loadContainers()
	if err != nil {
		return err
	}

	for _, container := range c.containers {
		err = container.Provision()
		if err != nil {
			return err
		}

	}
	c.log.DebugF("provisioned %d containers", len(c.containers))
	return nil
}

func (c *ContainersController) Start() error {

	c.log.TraceF("Starting containers")
	err := c.loadContainers()
	if err != nil {
		return err
	}

	for _, container := range c.containers {
		err = container.Start()
		if err != nil {
			return err
		}

	}
	c.log.DebugF("started %d containers", len(c.containers))
	return nil
}

func (c *ContainersController) Deploy() error {

	c.log.TraceF("Deploying containers")
	err := c.loadContainers()
	if err != nil {
		return err
	}

	for _, container := range c.containers {
		err = container.Deploy()
		if err != nil {
			return err
		}

	}
	c.log.DebugF("deployed %d containers", len(c.containers))
	return nil
}
func (c *ContainersController) Undeploy() error {

	c.log.TraceF("Undeploying containers")
	err := c.loadContainers()
	if err != nil {
		return err
	}

	for _, container := range c.containers {
		err = container.Undeploy()
		if err != nil {
			return err
		}

	}
	c.log.DebugF("undeployed %d containers", len(c.containers))
	return nil
}

func (c *ContainersController) Stop() error {

	plateformController, err := c.getPlateformController()
	if err != nil {
		return err
	}
	present, err := plateformController.IsPresent()
	if err != nil {
		return err
	}

	if present {
		err := plateformController.Stop()
		if err != nil {
			return err
		}

		c.log.TraceF("Stopping containers")
		err = c.loadContainers()
		if err != nil {
			return err
		}

		for _, container := range c.containers {
			status, err := container.Status()
			if err != nil {
				return err
			}

			if status != "Stopped" {

				err = container.Stop()
				if err != nil {
					return err
				}
				c.log.DebugF("stopped container %s", container.Name)
			}

		}

		c.log.DebugF("Stopped containers")
	}
	return nil
}

func (c *ContainersController) Rm() error {

	err := c.Stop()
	if err != nil {
		return err
	}

	plateformController, err := c.getPlateformController()
	if err != nil {
		return err
	}
	present, err := plateformController.IsPresent()
	if err != nil {
		return err
	}

	if present {
		c.log.TraceF("Remove containers")
		err = c.loadContainers()
		if err != nil {
			return err
		}

		for _, container := range c.containers {
			err = container.Unprovision()
			if err != nil {
				return err
			}

		}
		c.log.DebugF("removed %d containers", len(c.containers))
	}
	return nil
}

func (c *ContainersController) Unprovision() error {

	err := c.Rm()
	if err != nil {
		return err
	}
	plateformController, err := c.getPlateformController()
	if err != nil {
		return err
	}
	present, err := plateformController.IsPresent()
	if err != nil {
		return err
	}

	if present {

		c.log.TraceF("Unprovisioning containers")
		err = c.loadContainers()
		if err != nil {
			return err
		}

		for _, container := range c.containers {
			err = container.Unprovision()
			if err != nil {
				return err
			}

		}
		c.log.DebugF("unprovisioned %d containers", len(c.containers))
	}
	return nil
}
