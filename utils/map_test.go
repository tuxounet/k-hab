package utils

import (
	"testing"
)

func TestTtLoadYamlFromString(t *testing.T) {
	ctx := NewTestContext()
	yamlStr := `
name: "test"
`
	result := LoadYamlFromString[map[string]string](ctx, yamlStr)
	if result["name"] != "test" {
		t.Fatalf("Expected 'test', got '%s'", result["name"])
	}
}

func TestTTLoadJSONFromString(t *testing.T) {
	ctx := NewTestContext()
	jsonStr := `
{"name": "test"}
`
	result := LoadJSONFromString[map[string]string](ctx, jsonStr)
	if result["name"] != "test" {
		t.Fatalf("Expected 'test', got '%s'", result["name"])
	}
}

func TestTTGetMapValue(t *testing.T) {
	ctx := NewTestContext()
	m := map[string]interface{}{
		"name": "test",
	}
	result := GetMapValue(ctx, m, "name")
	if result != "test" {
		t.Fatalf("Expected 'test', got '%s'", result)
	}
}
