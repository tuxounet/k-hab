package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type CmdCall struct {
	Command string
	Args    []string
	Cwd     *string
}

func NewCmdCall(command string, args ...string) *CmdCall {
	return &CmdCall{
		Command: command,
		Args:    args,
		Cwd:     nil,
	}
}

func (c *CmdCall) String() string {

	out := c.Command + " "
	for _, arg := range c.Args {
		out += arg + " "
	}
	out = strings.TrimSpace(out)

	return out
}

func WithCmdCallBuilder(ctx *ScopeContext, habConfig map[string]interface{}, cmdPrefixKey string, cmdNameKey string, args ...string) *CmdCall {

	return ScopingWithReturn(ctx, "utils", "WithCmdCallBuilder", func(ctx *ScopeContext) *CmdCall {

		cmd_prefix := GetMapValue(ctx, habConfig, cmdPrefixKey).(string)
		cmd := GetMapValue(ctx, habConfig, cmdNameKey).(string)

		full := strings.TrimSpace(fmt.Sprintf("%s %s", cmd_prefix, cmd))

		segments := strings.Split(full, " ")
		segments = append(segments, args...)

		final := strings.TrimSpace(strings.Join(segments, " "))

		segments = strings.Split(final, " ")

		first := segments[0]
		rest := []string{}
		if len(segments) > 1 {
			rest = segments[1:]
		}

		return &CmdCall{
			Command: first,
			Args:    rest,
			Cwd:     nil,
		}
	})
}

func ExecSyncOutput(ctx *ScopeContext, query *CmdCall) error {

	return ctx.Scope("utils", "ExecSyncOutput", func(ctx *ScopeContext) {
		ctx.Log.DebugF("Executing command: %s", query.String())
		cmd := exec.Command(query.Command, query.Args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		if query.Cwd != nil {
			cmd.Dir = *query.Cwd
		}

		err := cmd.Start()
		ctx.Must(err)

		cmd.Wait()
		code := cmd.ProcessState.ExitCode()
		if code != 0 {
			ctx.Must(errors.New("command failed with exit code " + fmt.Sprint(code)))
		}
	})
}

func ExecSyncOutputMayFail(ctx *ScopeContext, query *CmdCall) int {
	return ScopingWithReturn(ctx, "utils", "ExecSyncOutputMayFail", func(ctx *ScopeContext) int {
		ctx.Log.DebugF("Executing command: %s", query.String())
		cmd := exec.Command(query.Command, query.Args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		if query.Cwd != nil {
			cmd.Dir = *query.Cwd
		}
		err := cmd.Start()
		ctx.Must(err)

		cmd.Wait()
		code := cmd.ProcessState.ExitCode()
		return code

	})
}

func CommandSyncOutput(ctx *ScopeContext, query *CmdCall) string {
	return ScopingWithReturn(ctx, "utils", "CommandSyncOutput", func(ctx *ScopeContext) string {
		ctx.Log.DebugF("Executing command: %s", query.String())
		cmd := exec.Command(query.Command, query.Args...)
		if query.Cwd != nil {
			cmd.Dir = *query.Cwd
		}
		out, err := cmd.Output()
		ctx.Must(err)
		return string(out)
	})
}

func CommandSyncJsonOutput(ctx *ScopeContext, query *CmdCall) map[string]interface{} {
	return ScopingWithReturn(ctx, "utils", "CommandSyncJsonOutput", func(ctx *ScopeContext) map[string]interface{} {
		ctx.Log.DebugF("Executing command: %s", query.String())
		cmd := exec.Command(query.Command, query.Args...)
		if query.Cwd != nil {
			cmd.Dir = *query.Cwd
		}
		out, err := cmd.Output()
		ctx.Must(err)

		var result map[string]interface{}
		ctx.Must(json.Unmarshal(out, &result))

		return result
	})

}

func CommandSyncJsonArrayOutput(ctx *ScopeContext, query *CmdCall) []map[string]interface{} {
	return ScopingWithReturn(ctx, "utils", "CommandSyncJsonArrayOutput", func(ctx *ScopeContext) []map[string]interface{} {
		ctx.Log.DebugF("Executing command: %s", query.String())
		cmd := exec.Command(query.Command, query.Args...)
		if query.Cwd != nil {
			cmd.Dir = *query.Cwd
		}
		out, err := cmd.Output()
		ctx.Must(err)

		var result []map[string]interface{}
		ctx.Must(json.Unmarshal(out, &result))

		return result
	})

}
