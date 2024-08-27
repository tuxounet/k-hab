package setup_test

import (
	"os"
	"testing"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/context"
	"github.com/tuxounet/k-hab/context/config"
	"github.com/tuxounet/k-hab/context/setup"
)

func TestTTSetup(t *testing.T) {

	ctx := context.NewTestContext(t)
	config := config.NewConfig(ctx.GetLogger(), map[string]string{
		"setup": "test",
	})

	setup := setup.NewSetup(ctx.GetLogger(), config, bases.SetupFile{
		Config:     config.GetCurrent(),
		Containers: []bases.SetupContainer{},
	})

	if setup == nil {
		t.Fatalf("Expected not nil, got nil")
	}

	err := setup.LoadDefaultSetup()
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

	if len(setup.SetupContainers) != 0 {
		t.Fatalf("Expected 0, got %d", len(setup.SetupContainers))
	}

	os.WriteFile("test.yaml", []byte("config:\n  a: b\ncontainers:\n  - name: test\n    image: test\n"), 0644)
	defer os.Remove("test.yaml")

	setup.LoadSetupFromYamlFile("test.yaml")

}

func TestTTSetupBadFile(t *testing.T) {
	ctx := context.NewTestContext(t)
	config := config.NewConfig(ctx.GetLogger(), map[string]string{
		"setup": "test",
	})

	setup := setup.NewSetup(ctx.GetLogger(), config, bases.SetupFile{
		Config:     config.GetCurrent(),
		Containers: []bases.SetupContainer{},
	})

	if setup == nil {
		t.Fatalf("Expected not nil, got nil")
	}

	err := setup.LoadSetupFromYamlFile("inexistant")
	if err == nil {

		t.Fatalf("Expected error, got nil")
	}
}

func TestTTSetupInvalidFile(t *testing.T) {
	ctx := context.NewTestContext(t)
	config := config.NewConfig(ctx.GetLogger(), map[string]string{
		"setup": "test",
	})

	setup := setup.NewSetup(ctx.GetLogger(), config, bases.SetupFile{
		Config:     config.GetCurrent(),
		Containers: []bases.SetupContainer{},
	})

	if setup == nil {
		t.Fatalf("Expected not nil, got nil")
	}

	os.WriteFile("invalid.yaml", []byte("config: a:\n"), 0644)
	defer os.Remove("invalid.yaml")

	err := setup.LoadSetupFromYamlFile("invalid.yaml")
	if err == nil {

		t.Fatalf("Expected error, got nil")
	}
}
