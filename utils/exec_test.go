package utils

import "testing"

func TestTTCmd(t *testing.T) {

	cmd := "echo"
	args := []string{"hello"}
	result := NewCmdCall(cmd, args...)
	cwd := "plop"
	result.Cwd = &cwd
	if result.Command != "echo" {
		t.Fatalf("Expected 'echo', got '%s'", result.Command)
	}
	if result.Args[0] != "hello" {
		t.Fatalf("Expected 'hello', got '%s'", result.Args[0])
	}
	if *result.Cwd != "plop" {
		t.Fatalf("Expected 'plop', got '%s'", *result.Cwd)
	}
	str := result.String()
	if str != "echo hello" {
		t.Fatalf("Expected 'echo hello', got '%s'", str)
	}

}
