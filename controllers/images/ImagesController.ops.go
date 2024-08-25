package images

import (
	"errors"

	"github.com/tuxounet/k-hab/controllers/images/definitions"
)

func (h *ImagesController) loadImages() error {

	if h.images != nil {
		return nil
	}

	for _, confContainer := range h.ctx.GetSetupContainers() {

		found := false
		for _, localBase := range h.images {
			if localBase.Name == confContainer.Base {
				found = true
				break
			}
		}
		if !found {
			confImage, err := definitions.GetImageBase(confContainer.Base)
			if err != nil {
				return err
			}
			image := NewImageModel(confContainer.Base, h.ctx, confImage)
			h.images = append(h.images, image)
		}
	}
	return nil
}

func (h *ImagesController) GetImage(name string) (*ImageModel, error) {

	err := h.loadImages()
	if err != nil {
		return nil, err
	}
	for _, image := range h.images {
		if image.Name == name {
			return image, nil
		}
	}
	return nil, errors.New("image not found")

}

func (h *ImagesController) ImagePresent(name string) (bool, error) {

	err := h.loadImages()
	if err != nil {
		return false, err
	}
	image, err := h.GetImage(name)
	if err != nil {
		return false, err
	}
	return image.present()

}

func (h *ImagesController) EnsureImage(name string) (bool, error) {
	changed := false
	h.log.TraceF("Ensuring image %s", name)
	definition, err := definitions.GetImageBase(name)
	if err != nil {
		return false, err
	}

	present, err := h.ImagePresent(name)
	if err != nil {
		return false, err
	}

	if !present {
		h.log.WarnF("Image %s not present, provisioning", name)

		image := NewImageModel(name, h.ctx, definition)
		err = image.provision()
		if err != nil {
			return false, err
		}
		h.images = append(h.images, image)
		h.log.DebugF("Image %s provisioned", name)
		changed = true
	}

	image, err := h.GetImage(name)
	if err != nil {
		return false, err
	}

	needBuild, err := image.needBuild(definition)
	if err != nil {
		return false, err
	}
	if needBuild {
		h.log.WarnF("Image %s need rebuild, provisioning", name)

		err = image.provision()
		if err != nil {
			return false, err
		}
		h.log.DebugF("Image %s provisioned", name)

	}

	return changed, err

}
