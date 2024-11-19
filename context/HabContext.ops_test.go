package context

import (
	"testing"
)

func TestHabContextPropsSetup(t *testing.T) {
	ctx := NewTestContext(t)

	_, err := ctx.getEntryContainer()
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
	err = ctx.Init()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}

	err = ctx.SetSetup("")
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	container, err := ctx.getEntryContainer()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	if container == nil {
		t.Errorf("Expected not nil but got nil")
	}

}
