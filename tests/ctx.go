package tests

import (
	ctx "context"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/context"
)

func NewTestContext() bases.IContext {
	rootCtx := ctx.TODO()
	return context.NewHabContext(rootCtx)
}
