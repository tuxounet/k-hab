package context

import (
	"context"

	"github.com/tuxounet/k-hab/bases"
)

func NewTestContext() bases.IContext {

	return NewHabContext(context.TODO(), map[string]string{}, bases.SetupFile{})

}
