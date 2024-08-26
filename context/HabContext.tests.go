package context

import (
	"context"
	"log"
	"os"

	"github.com/tuxounet/k-hab/bases"
)

func NewTestContext() *HabContext {
	workingFolder, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return NewHabContext(context.TODO(), map[string]string{}, bases.SetupFile{}, workingFolder)

}
