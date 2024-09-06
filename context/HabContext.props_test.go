package context_test

import (
	"os"
	"testing"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/context"
)

func TestHabContextProps(t *testing.T) {
	ctx := context.NewTestContext(t)
	currentConfig := ctx.GetCurrentConfig()
	ctx.SetConfigValue("test", "test")
	if currentConfig["test"] != "test" {
		t.Errorf("Expected 'test' but got '%s'", currentConfig["test"])
	}

	setupContainers := ctx.GetSetupContainers()
	if len(setupContainers) != 0 {
		t.Errorf("Expected 0 but got %d", len(setupContainers))
	}

	notFoundController, err := ctx.GetController(bases.BuilderController)
	if err != nil {
		if err.Error() != "controller not found" {
			t.Errorf("Expected 'controller not found' but got '%s'", err.Error())
		}

	}
	if notFoundController != nil {
		t.Errorf("Expected nil but got %v", notFoundController)
	}

	err = ctx.Init()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}

	stroot, err := ctx.GetStorageRoot()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	if stroot == "" {
		t.Errorf("Expected not empty but got empty")
	}

	ctx.SetConfigValue("hab.storage.root", "/tmp")
	stroot, err = ctx.GetStorageRoot()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	if stroot != "/tmp" {
		t.Errorf("Expected '/tmp' but got '%s'", stroot)
	}
	containers := ctx.GetSetupContainers()
	if len(containers) != 0 {
		t.Errorf("Expected 0 but got %d", len(containers))
	}

	err = ctx.SetSetup("")
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	conts := ctx.GetSetupContainers()
	if len(conts) == 0 {
		t.Errorf("Expected not 0 but got %d", len(conts))
	}

}

func TestHabContextPropsSetup(t *testing.T) {
	ctx := context.NewTestContext(t)
	setupBody := `config:[]\ncontainers: []`
	os.WriteFile("test.yaml", []byte(setupBody), 0644)
	defer os.Remove("test.yaml")

	err := ctx.SetSetup("test.yaml")
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	conts := ctx.GetSetupContainers()
	if len(conts) != 0 {
		t.Errorf("Expected 0 but got %d", len(conts))
	}

}

func TestHabContextPropsLevel(t *testing.T) {
	ctx := context.NewTestContext(t)

	err := ctx.SetLogLevel("DEBUG")
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	err = ctx.SetLogLevel("TRACE")
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	err = ctx.SetLogLevel("INFO")
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	err = ctx.SetLogLevel("ERROR")
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}
	err = ctx.SetLogLevel("WARN")
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}

	err = ctx.SetLogLevel("FATAL")
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}

	err = ctx.SetLogLevel("UNKNOWN")
	if err == nil {
		t.Errorf("Expected error but got nil")
	}

}
