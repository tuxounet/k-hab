package utils

import (
	"os"
	"testing"
)

func TestTTCmdCall(t *testing.T) {

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

func TestTTJsonOutput(t *testing.T) {
	ctx := NewTestContext()
	cmd := NewCmdCall("echo", `{"hello": "world"}`)
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory")
	}
	cmd.Cwd = &cwd

	out := JsonCommandOutput[map[string]string](ctx, cmd)
	if out["hello"] != "world" {
		t.Fatalf("Expected 'world', got '%s'", out["hello"])
	}
}

func TestTTExecExitCode(t *testing.T) {
	ctx := NewTestContext()
	cmd := NewCmdCall("ls")
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory")
	}
	cmd.Cwd = &cwd
	code := OsExecWithExitCode(ctx, cmd)
	if code != 0 {
		t.Fatalf("Expected 0, got %d", code)
	}

}
func TestTTOSExec(t *testing.T) {
	ctx := NewTestContext()
	cmd := NewCmdCall("ls")
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory")
	}
	cmd.Cwd = &cwd
	err = OsExec(ctx, cmd)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

}
func TestTTFail(t *testing.T) {
	ctx := NewTestContext()
	cmd := NewCmdCall("unexistant", "command")
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory")
	}
	cmd.Cwd = &cwd

	err = OsExec(ctx, cmd)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}

}

func TestTTCmdBuilder(t *testing.T) {
	ctx := NewTestContext()
	const jsonString = `{"cmd": { "prefix" : "", "name" : "ls" }}`
	habConfig := LoadJSONFromString[map[string]interface{}](ctx, jsonString)

	cmd := WithCmdCall(ctx, habConfig, "cmd.prefix", "cmd.name", "-l")
	if cmd.Command != "ls" {
		t.Fatalf("Expected 'ls', got '%s'", cmd.Command)
	}
	if cmd.Args[0] != "-l" {
		t.Fatalf("Expected '-l', got '%s'", cmd.Args[0])
	}

}
