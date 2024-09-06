package images

import (
	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/controllers/runtime"
)

func (h *ImagesController) Provision() error {
	h.log.TraceF("Provisioning")
	err := h.loadImages()
	if err != nil {
		return err
	}

	h.log.DebugF("Provisioned")

	return nil

}

func (h *ImagesController) Unprovision() error {
	h.log.TraceF("Unprovisioning")
	controller, err := h.ctx.GetController(bases.RuntimeController)
	if err != nil {
		return err
	}
	runtimeController := controller.(*runtime.RuntimeController)
	present, err := runtimeController.IsPresent()
	if err != nil {
		return err
	}

	if present {

		err := h.loadImages()
		if err != nil {
			return err
		}

		for _, image := range h.images {
			err := image.unprovision()
			if err != nil {
				return err
			}
		}
		h.log.DebugF("Unprovisionned %d images", len(h.images))
	}
	return nil
}

func (h *ImagesController) Nuke() error {
	h.log.TraceF("Nuking")
	err := h.loadImages()
	if err != nil {
		return nil
	}

	for _, image := range h.images {
		err := image.nuke()
		if err != nil {
			return err
		}
	}
	h.log.DebugF("Nuked %d images", len(h.images))
	return nil
}
