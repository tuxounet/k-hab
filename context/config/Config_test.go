package config_test

import (
	"context"
	"testing"

	"github.com/tuxounet/k-hab/context/config"
	"github.com/tuxounet/k-hab/context/logger"
)

func TestTTDefaultConfig(t *testing.T) {

	log := logger.NewLogger(context.TODO(), "TEST")
	config := config.NewConfig(log, map[string]string{
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

	log := logger.NewLogger(context.TODO(), "TEST")
	config := config.NewConfig(log, map[string]string{})

	err := config.SetConfigValue("a.b.c", "1")
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if config.GetValue("a.b.c") != "1" {
		t.Fatalf("Expected '1', got '%s'", config.GetValue("a.b.c"))
	}
}
