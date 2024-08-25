package images

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