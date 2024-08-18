package hab

import "github.com/tuxounet/k-hab/utils"

func (h *Hab) loadImages(ctx *utils.ScopeContext) error {
	return ctx.Scope(h.scopeBase, "loadImages", func(ctx *utils.ScopeContext) {
		if h.images != nil {
			return
		}
		for _, confImage := range h.config.ImagesConfig {

			found := false
			for _, localImage := range h.images {
				if localImage.name == confImage.Name {
					found = true
					break
				}
			}
			if !found {
				image := newHabImage(confImage.Name, h, &confImage)
				h.images = append(h.images, image)
			}
		}
	})
}

func (h *Hab) getImage(ctx *utils.ScopeContext, name string) *HabImage {
	return utils.ScopingWithReturnOnly(ctx, h.scopeBase, "getImage", func(ctx *utils.ScopeContext) *HabImage {
		ctx.Must(h.loadImages(ctx))
		for _, image := range h.images {
			if image.name == name {
				return image
			}
		}
		ctx.Must(ctx.Error("Image not found"))
		return nil
	})
}

func (h *Hab) imagePresent(ctx *utils.ScopeContext, name string) bool {
	return utils.ScopingWithReturnOnly(ctx, h.scopeBase, "imagePresent", func(ctx *utils.ScopeContext) bool {
		ctx.Must(h.loadImages(ctx))
		image := h.getImage(ctx, name)
		return image.present(ctx)

	})
}

func (h *Hab) provisionImages(ctx *utils.ScopeContext) error {
	return ctx.Scope(h.scopeBase, "provisionImages", func(ctx *utils.ScopeContext) {
		ctx.Must(h.builder.Provision(ctx))
		ctx.Must(h.loadImages(ctx))

		for _, image := range h.images {
			if !image.present(ctx) {
				ctx.Must(image.provision(ctx))
			}

		}
		ctx.Log.InfoF("Provisioned %d images", len(h.images))

	})
}

func (h *Hab) removeImages(ctx *utils.ScopeContext) error {
	return ctx.Scope(h.scopeBase, "provisionImages", func(ctx *utils.ScopeContext) {
		ctx.Must(h.builder.Provision(ctx))
		ctx.Must(h.loadImages(ctx))

		for _, image := range h.images {
			ctx.Must(image.unprovision(ctx))
		}
		ctx.Log.InfoF("Provisioned %d images", len(h.images))

	})
}

func (h *Hab) unprovisionImages(ctx *utils.ScopeContext) error {
	return ctx.Scope(h.scopeBase, "unprovisionImages", func(ctx *utils.ScopeContext) {
		ctx.Must(h.loadImages(ctx))

		for _, image := range h.images {
			ctx.Must(image.unprovision(ctx))
		}
	})
}
