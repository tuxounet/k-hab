package utils_test

import (
	"os"
	"testing"

	"github.com/tuxounet/k-hab/context"
	"github.com/tuxounet/k-hab/utils"
)

func TestTTCmdCall(t *testing.T) {

	cmd := "echo"
	args := []string{"hello"}
	result := utils.NewCmdCall(cmd, args...)
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

	cmd := utils.NewCmdCall("echo", `{"hello": "world"}`)
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory")
	}
	cmd.Cwd = &cwd

	out, err := utils.JsonCommandOutput[map[string]string](cmd)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if out["hello"] != "world" {
		t.Fatalf("Expected 'world', got '%s'", out["hello"])
	}
}

func TestTTJsonOutputFailed(t *testing.T) {

	cmd := utils.NewCmdCall("inexistant", `{"hello": "world"}`)

	_, err := utils.JsonCommandOutput[map[string]string](cmd)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}

}

func TestTTInvalidJsonOutput(t *testing.T) {

	cmd := utils.NewCmdCall("echo", `{"hello": "worl"`)
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory")
	}
	cmd.Cwd = &cwd

	_, err = utils.JsonCommandOutput[map[string]string](cmd)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}

}

func TestTTExecExitCode(t *testing.T) {

	cmd := utils.NewCmdCall("ls")
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory")
	}
	cmd.Cwd = &cwd
	code, err := utils.OsExecWithExitCode(cmd)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if code != 0 {
		t.Fatalf("Expected 0, got %d", code)
	}

}
func TestTTOSExec(t *testing.T) {

	cmd := utils.NewCmdCall("ls")
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory")
	}
	cmd.Cwd = &cwd
	err = utils.OsExec(cmd)
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}

}
func TestTTFail(t *testing.T) {

	cmd := utils.NewCmdCall("unexistant", "command")
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory")
	}
	cmd.Cwd = &cwd

	err = utils.OsExec(cmd)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}

}
func TestTTINonZero(t *testing.T) {

	cmd := utils.NewCmdCall("which", "lsd")
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Error getting current working directory")
	}
	cmd.Cwd = &cwd

	err = utils.OsExec(cmd)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}

}

func TestTTCmdBuilder(t *testing.T) {

	ctx := context.NewTestContext()
	ctx.SetConfigValue("cmd.prefix", "")
	ctx.SetConfigValue("cmd.name", "ls")

	cmd, err := utils.WithCmdCall(ctx, "cmd.prefix", "cmd.name", "-l")
	if err != nil {
		t.Fatalf("Error building command: %v", err)
	}
	if cmd.Command != "ls" {
		t.Fatalf("Expected 'ls', got '%s'", cmd.Command)
	}
	if cmd.Args[0] != "-l" {
		t.Fatalf("Expected '-l', got '%s'", cmd.Args[0])
	}

}
