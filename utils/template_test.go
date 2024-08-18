package utils

import "testing"

func TestUntemplate(t *testing.T) {
	ctx := NewTestContext()
	tpl := "Hello {{.Name}}"
	data := struct {
		Name string
	}{
		Name: "World",
	}
	result := UnTemplate(ctx, tpl, data)
	if result != "Hello World" {
		t.Fatalf("Expected 'Hello World', got '%s'", result)
	}
}
