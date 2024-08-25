package utils_test

import (
	"testing"

	"github.com/tuxounet/k-hab/utils"
)

func TestUntemplate(t *testing.T) {
	tpl := "Hello {{.Name}}"
	data := struct {
		Name string
	}{
		Name: "World",
	}
	result, err := utils.UnTemplate(tpl, data)
	if err != nil {
		t.Fatalf("Error untemplating: %v", err)
	}
	if result != "Hello World" {
		t.Fatalf("Expected 'Hello World', got '%s'", result)
	}
}

func TestUntemplateInvalidTemplate(t *testing.T) {
	tpl := "Hello {{ UNDEfined }}"
	data := struct {
		Name string
	}{
		Name: "World",
	}
	result, err := utils.UnTemplate(tpl, data)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	if result != "" {
		t.Fatalf("Expected '', got '%s'", result)
	}
}

func TestUntemplateBadTemplate(t *testing.T) {
	tpl := "Hello {{.UNDEfined}}"
	data := struct {
		Name string
	}{
		Name: "World",
	}
	result, err := utils.UnTemplate(tpl, data)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
	if result != "" {
		t.Fatalf("Expected '', got '%s'", result)
	}
}
