package images

import (
	"github.com/tuxounet/k-hab/bases"
)

type ImagesController struct {
	bases.BaseController
	ctx    bases.IContext
	log    bases.ILogger
	images []*ImageModel
}

func NewImagesController(ctx bases.IContext) *ImagesController {
	return &ImagesController{
		ctx: ctx,
		log: ctx.GetSubLogger("ImagesController", ctx.GetLogger()),
	}
}

func (h *ImagesController) Provision() error {
	h.log.TraceF("Provisioning")
	err := h.loadImages()
	if err != nil {
		return err
	}

	for _, image := range h.images {
		present, err := h.ImagePresent(image.name)
		if err != nil {
			return err
		}

		if !present {

			err := image.provision()
			if err != nil {
				return err
			}
		}

	}
	h.log.DebugF("Provisioned %d images", len(h.images))

	return nil

}

func (h *ImagesController) Unprovision() error {
	h.log.TraceF("Unprovisioning")
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
	h.log.DebugF("Unprovisioned %d images", len(h.images))
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
