package containers

import (
	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/controllers/runtime"
)

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
		err = container.Start()
		if err != nil {
			return err
		}

	}
	c.log.DebugF("deployed %d containers", len(c.containers))
	return nil
}

func (c *ContainersController) Stop() error {

	controller, err := c.ctx.GetController(bases.RuntimeController)
	if err != nil {
		return err
	}
	runtimeController := controller.(*runtime.RuntimeController)
	present, err := runtimeController.IsPresent()
	if err != nil {
		return err
	}

	if present {
		err := runtimeController.Stop()
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
	return nil
}

func (c *ContainersController) Unprovision() error {

	err := c.Rm()
	if err != nil {
		return err
	}

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
	return nil
}
