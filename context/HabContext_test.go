package context_test

import (
	"testing"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/context"
)

func TestBaseHabContext(t *testing.T) {
	ctx := context.NewTestContext(t)
	ctx.SetConfigValue("test", "test")

	value := ctx.GetConfigValue("test")
	if value != "test" {
		t.Errorf("Expected 'test' but got '%s'", value)
	}

}

func TestHabContextLogger(t *testing.T) {
	ctx := context.NewTestContext(t)
	logger := ctx.GetLogger()

	logger.InfoF("Test")
	sub := ctx.GetSubLogger("test", logger)
	sub.InfoF("Test")
}

func TestTTHabControllers(t *testing.T) {
	ctx := context.NewTestContext(t)
	err := ctx.Init()
	if err != nil {
		t.Errorf("Expected nil but got %v", err)
	}

	order := bases.HabControllersLoadOrder()

	for _, controllerType := range order {
		controller, err := ctx.GetController(controllerType)
		if err != nil {
			t.Errorf("Expected nil but got %v", err)
		}
		if controller == nil {
			t.Errorf("Expected not nil but got nil")
		}

	}

	_, err = ctx.GetController(bases.HabControllers("test"))
	if err == nil {
		t.Errorf("Expected error but got nil")
	}

}
