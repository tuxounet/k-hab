package context_test

import (
	"testing"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/context"
)

func TestBaseHabContext(t *testing.T) {
	ctx := context.NewTestContext()
	ctx.SetConfigValue("test", "test")

	value := ctx.GetConfigValue("test")
	if value != "test" {
		t.Errorf("Expected 'test' but got '%s'", value)
	}

	currentConfig := ctx.GetCurrentConfig()
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

}

func TestHabContextLogger(t *testing.T) {
	ctx := context.NewTestContext()
	logger := ctx.GetLogger()

	logger.InfoF("Test")
	sub := ctx.GetSubLogger("test", logger)
	sub.InfoF("Test")
}

// func TestTTHabLifecycle(t *testing.T) {
// 	ctx := context.NewTestContext()

// 	err := ctx.Provision()
// 	if err != nil {
// 		t.Errorf("Expected nil but got %v", err)
// 	}

// 	err = ctx.Unprovision()
// 	if err != nil {
// 		t.Errorf("Expected nil but got %v", err)
// 	}

// }
