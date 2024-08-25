package images

import (
	"errors"

	"github.com/tuxounet/k-hab/controllers/images/definitions"
)

func (h *ImagesController) loadImages() error {

	if h.images != nil {
		return nil
	}

	for _, confContainer := range h.ctx.GetContainersConfig() {

		found := false
		for _, localBase := range h.images {
			if localBase.name == confContainer.Base {
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
		if image.name == name {
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

func (h *ImagesController) EnsureImage(name string) error {

	return nil

}
