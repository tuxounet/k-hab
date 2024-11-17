package context

import (
	"github.com/tuxounet/k-hab/bases"
)

func (h *HabContext) Install() error {
	h.log.DebugF("Hab installing...")
	//Start
	for _, controllerKey := range bases.HabControllersLoadOrder() {
		controller, err := h.GetController(controllerKey)
		if err != nil {
			return err
		}
		err = controller.Install()
		if err != nil {
			return err
		}
	}

	h.log.InfoF("Hab Installed")
	return nil
}

func (h *HabContext) Uninstall() error {

	h.log.DebugF("Hab uninstalling...")
	//Start
	for _, controllerKey := range bases.HabControllersUnloadOrder() {
		controller, err := h.GetController(controllerKey)
		if err != nil {
			return err
		}
		err = controller.Uninstall()
		if err != nil {
			return err
		}
	}
	h.log.InfoF("Hab Uninstalled")

	return nil

}

func (h *HabContext) Provision() error {
	h.log.DebugF("Hab provisionning...")

	for _, controllerKey := range bases.HabControllersLoadOrder() {
		controller, err := h.GetController(controllerKey)
		if err != nil {
			return err
		}

		err = controller.Install()
		if err != nil {
			return err
		}

		err = controller.Provision()
		if err != nil {
			return err
		}
	}
	h.log.InfoF("Hab Provisioned")
	return nil
}

func (h *HabContext) Start() error {

	//Ensure Provisioning
	err := h.Provision()
	if err != nil {
		return err
	}

	h.log.DebugF("Hab starting...")
	//Start
	for _, controllerKey := range bases.HabControllersLoadOrder() {
		controller, err := h.GetController(controllerKey)
		if err != nil {
			return err
		}
		err = controller.Start()
		if err != nil {
			return err
		}
	}

	h.log.InfoF("Hab Started")

	return nil

}
func (h *HabContext) Deploy() error {

	//Ensure Provisioning
	err := h.Start()
	if err != nil {
		return err
	}
	h.log.DebugF("Hab Deploying...")
	//Start
	for _, controllerKey := range bases.HabControllersLoadOrder() {
		controller, err := h.GetController(controllerKey)
		if err != nil {
			return err
		}
		err = controller.Deploy()
		if err != nil {
			return err
		}
	}
	h.log.InfoF("Hab Deployed")

	return nil

}

func (h *HabContext) Shell() error {

	//Ensure Start
	err := h.Start()
	if err != nil {
		return err
	}

	h.log.DebugF("Looking for an entry container")
	container, err := h.getEntryContainer()
	if err != nil {
		return err
	}

	h.log.DebugF("Waiting for container to be ready")
	err = container.WaitReady()
	if err != nil {
		return err
	}

	h.log.DebugF("Starting shell")

	err = container.Shell()
	if err != nil {
		return err
	}

	h.log.InfoF("Hab Shell completed")
	return nil

}

func (h *HabContext) Stop() error {

	h.log.TraceF("Hab Stopping...")
	for _, controllerKey := range bases.HabControllersUnloadOrder() {
		controller, err := h.GetController(controllerKey)
		if err != nil {
			return err
		}

		err = controller.Stop()
		if err != nil {
			return err
		}
		h.log.DebugF("Controller %s stopped", controllerKey)
	}
	h.log.InfoF("Hab Stopped")
	return nil

}
func (h *HabContext) Undeploy() error {

	h.log.DebugF("Hab Undeploying...")
	//Start
	for _, controllerKey := range bases.HabControllersLoadOrder() {
		controller, err := h.GetController(controllerKey)
		if err != nil {
			return err
		}
		err = controller.Undeploy()
		if err != nil {
			return err
		}
	}
	h.log.InfoF("Hab Undeployed")

	return nil

}
func (h *HabContext) Rm() error {
	err := h.Stop()
	if err != nil {
		return err
	}

	h.log.DebugF("Hab Removing...")
	for _, controllerKey := range bases.HabControllersUnloadOrder() {
		controller, err := h.GetController(controllerKey)
		if err != nil {
			return err
		}
		err = controller.Rm()
		if err != nil {
			return err
		}
		h.log.DebugF("Controller %s removed", controllerKey)
	}
	h.log.InfoF("Hab Removed")
	return nil

}

func (h *HabContext) Unprovision() error {

	err := h.Rm()
	if err != nil {
		return err
	}
	h.log.DebugF("Hab Unprovisioning...")
	for _, controllerKey := range bases.HabControllersUnloadOrder() {
		controller, err := h.GetController(controllerKey)
		if err != nil {
			return err
		}
		err = controller.Unprovision()
		if err != nil {
			return err
		}
		h.log.DebugF("Controller %s unprovisioned", controllerKey)
	}
	h.log.InfoF("Hab Unprovisioned")
	return nil

}

func (h *HabContext) Nuke() error {
	err := h.Uninstall()
	if err != nil {
		return err
	}
	h.log.DebugF("Hab Nuking...")
	for _, controllerKey := range bases.HabControllersUnloadOrder() {
		controller, err := h.GetController(controllerKey)
		if err != nil {
			return err
		}
		err = controller.Nuke()
		if err != nil {
			return err
		}
	}
	h.log.InfoF("Hab Nuked")
	return nil

}
