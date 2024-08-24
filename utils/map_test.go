package utils

import (
	"testing"
)

func TestTtLoadYamlFromString(t *testing.T) {

	yamlStr := `
name: "test"
`
	result, err := LoadYamlFromString[map[string]string](yamlStr)
	if err != nil {
		t.Fatalf("Error loading yaml: %v", err)
	}
	if result["name"] != "test" {
		t.Fatalf("Expected 'test', got '%s'", result["name"])
	}
}

func TestTTLoadJSONFromString(t *testing.T) {

	jsonStr := `
{"name": "test"}
`
	result, err := LoadJSONFromString[map[string]string](jsonStr)
	if err != nil {
		t.Fatalf("Error loading json: %v", err)
	}
	if result["name"] != "test" {
		t.Fatalf("Expected 'test', got '%s'", result["name"])
	}
}

func TestTTGetMapValue(t *testing.T) {

	m := map[string]interface{}{
		"name": "test",
	}
	result := GetMapValue(m, "name")

	if result != "test" {
		t.Fatalf("Expected 'test', got '%s'", result)
	}
}
