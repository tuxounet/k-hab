package images

import (
	"errors"
)

func (h *ImagesController) loadImages() error {

	if h.images != nil {
		return nil
	}
	for _, confImage := range h.ctx.GetImagesConfig() {

		found := false
		for _, localImage := range h.images {
			if localImage.name == confImage.Name {
				found = true
				break
			}
		}
		if !found {
			image := NewImageModel(confImage.Name, h.ctx, confImage)
			h.images = append(h.images, image)
		}
	}
	return nil
}

func (h *ImagesController) getImage(name string) (*ImageModel, error) {

	err := h.loadImages()
	if err != nil {
		return nil, err
	}
	for _, image := range h.images {
		if image.name == name {
			return image, nil
		}
	}
	return nil, errors.New("image not found")

}

func (h *ImagesController) imagePresent(name string) (bool, error) {

	err := h.loadImages()
	if err != nil {
		return false, err
	}
	image, err := h.getImage(name)
	if err != nil {
		return false, err
	}
	return image.present()

}
