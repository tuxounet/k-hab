package config_test

import (
	"testing"

	"github.com/tuxounet/k-hab/context/config"
)

func TestTTDefaultConfig(t *testing.T) {

	config := config.NewConfig(map[string]string{
		"a.b.c": "1",
	})

	if config.GetValue("a.b.c") != "1" {
		t.Fatalf("Expected '1', got '%s'", config.GetValue("a.b.c"))
	}

	current := config.GetCurrent()
	if current["a.b.c"] != "1" {
		t.Fatalf("Expected '1', got '%s'", current["a.b.c"])
	}
}

func TestTTSetConfigValue(t *testing.T) {

	config := config.NewConfig(map[string]string{})

	err := config.SetConfigValue("a.b.c", "1")
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if config.GetValue("a.b.c") != "1" {
		t.Fatalf("Expected '1', got '%s'", config.GetValue("a.b.c"))
	}
}
