package config

import (
	"testing"

	"github.com/tuxounet/k-hab/utils"
)

func TestTTDefaultConfig(t *testing.T) {
	ctx := utils.NewTestContext()
	config := NewConfig()
	err := config.Load(ctx)
	if err != nil {
		t.Fatalf("Error loading config: %s", err)
	}
	if config.HabConfig == nil {
		t.Fatalf("HabConfig is nil")
	}
	if len(config.ContainersConfig) == 0 {
		t.Fatalf("ContainersConfig is empty")
	}
	if len(config.ImagesConfig) == 0 {
		t.Fatalf("ImagesConfig is empty")
	}

	containerConfig := config.GetContainerConfig(ctx, "bastion")
	if containerConfig.Name != "bastion" {
		t.Fatalf("ContainerConfig.Name is not bastion")
	}

	// expect panic
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Recovered from panic: %s", r)
		}
	}()

	config.GetContainerConfig(ctx, "non-existing")

}
