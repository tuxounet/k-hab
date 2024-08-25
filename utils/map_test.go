package utils_test

import (
	"testing"

	"github.com/tuxounet/k-hab/utils"
)

func TestTtLoadYamlFromString(t *testing.T) {

	yamlStr := `
name: "test"
`
	result, err := utils.LoadYamlFromString[map[string]string](yamlStr)
	if err != nil {
		t.Fatalf("Error loading yaml: %v", err)
	}
	if result["name"] != "test" {
		t.Fatalf("Expected 'test', got '%s'", result["name"])
	}
}

func TestTTLoadYamlFromInvalidString(t *testing.T) {

	yamlStr := `
name:  est"
  jdj: 1
`
	_, err := utils.LoadYamlFromString[map[string]string](yamlStr)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestTTLoadJSONFromString(t *testing.T) {

	jsonStr := `
{"name": "test"}
`
	result, err := utils.LoadJSONFromString[map[string]string](jsonStr)
	if err != nil {
		t.Fatalf("Error loading json: %v", err)
	}
	if result["name"] != "test" {
		t.Fatalf("Expected 'test', got '%s'", result["name"])
	}
}
func TestTTLoadJSONFromInvalidString(t *testing.T) {

	jsonStr := `
{"name": "test
`
	result, err := utils.LoadJSONFromString[map[string]string](jsonStr)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}

	if result != nil {
		t.Fatalf("Expected nil, got '%v'", result)
	}
}
