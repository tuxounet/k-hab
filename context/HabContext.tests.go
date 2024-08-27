package context

import (
	"context"
	"os"
	"testing"
)

func NewTestContext(t *testing.T) *HabContext {
	workingFolder, _ := os.Getwd()

	return NewHabContext(context.TODO(), workingFolder)

}
