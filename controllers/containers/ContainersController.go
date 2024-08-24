package containers

import (
	"github.com/tuxounet/k-hab/bases"
)

type ContainersController struct {
	bases.BaseController
	ctx        bases.IContext
	log        bases.ILogger
	containers map[string]ContainerModel
}

func NewContainersController(ctx bases.IContext) *ContainersController {

	return &ContainersController{
		ctx:        ctx,
		log:        ctx.GetSubLogger("ContainersController", ctx.GetLogger()),
		containers: make(map[string]ContainerModel),
	}
}

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

func (c *ContainersController) Stop() error {

	c.log.TraceF("Stopping containers")
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
	c.log.DebugF("stopped %d containers", len(c.containers))
	return nil
}

func (c *ContainersController) Unprovision() error {

	c.log.TraceF("Unprovisioning containers")
	err := c.loadContainers()
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
