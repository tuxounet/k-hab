package config

import (
	"testing"
)

func TestTTDefaultConfig(t *testing.T) {

	config := NewConfig()
	err := config.Load()
	if err != nil {
		t.Fatalf("Error loading config: %s", err)
	}
	if config.HabConfig == nil {
		t.Fatalf("HabConfig is nil")
	}
	if len(config.ContainersConfig) == 0 {
		t.Fatalf("ContainersConfig is empty")
	}

	containerConfig, err := config.GetContainerConfig("bastion")
	if err != nil {
		t.Fatalf("Error getting container config: %s", err)
	}

	if containerConfig.Name != "bastion" {
		t.Fatalf("ContainerConfig.Name is not bastion")
	}

	// expect panic
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Recovered from panic: %s", r)
		}
	}()

	_, err = config.GetContainerConfig("non-existing")
	if err == nil {
		t.Fatalf("Expected error")
	}

}
