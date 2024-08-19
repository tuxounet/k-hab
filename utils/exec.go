package utils

import (
	"encoding/json"
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

func WithCmdCall(ctx *ScopeContext, anyMap any, cmdPrefixKey string, cmdNameKey string, args ...string) *CmdCall {

	return ScopingWithReturn(ctx, "utils", "WithCmdCallBuilder", func(ctx *ScopeContext) *CmdCall {

		cmd_prefix := GetMapValue(ctx, anyMap, cmdPrefixKey).(string)
		cmd := GetMapValue(ctx, anyMap, cmdNameKey).(string)

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

func OsExec(ctx *ScopeContext, query *CmdCall) error {

	return ScopingWithReturn(ctx, "utils", "OsExec", func(ctx *ScopeContext) error {
		code := OsExecWithExitCode(ctx, query)
		if code != 0 {
			return ctx.Error("command failed with exit code " + fmt.Sprint(code))
		}
		return nil
	})
}

func OsExecWithExitCode(ctx *ScopeContext, query *CmdCall) int {
	return ScopingWithReturn(ctx, "utils", "OsExecWithExitCode", func(ctx *ScopeContext) int {
		ctx.Log.DebugF("Executing command: %s", query.String())
		cmd := exec.Command(query.Command, query.Args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		if query.Cwd != nil {
			cmd.Dir = *query.Cwd
		}

		err := cmd.Start()
		if err != nil {
			ctx.Log.WarnF("Error starting command: %s", err)
			return -1
		}

		cmd.Wait()
		return cmd.ProcessState.ExitCode()

	})

}

func JsonCommandOutput[R any](ctx *ScopeContext, query *CmdCall) R {
	return ScopingWithReturn(ctx, "utils", "JsonCommandOutput", func(ctx *ScopeContext) R {
		out := RawCommandOutput(ctx, query)

		var result R
		ctx.Must(json.Unmarshal([]byte(out), &result))
		return result
	})
}

func RawCommandOutput(ctx *ScopeContext, query *CmdCall) string {
	return ScopingWithReturn(ctx, "utils", "RawCommandOutput", func(ctx *ScopeContext) string {
		ctx.Log.DebugF("Executing command: %s", query.String())
		cmd := exec.Command(query.Command, query.Args...)
		if query.Cwd != nil {
			cmd.Dir = *query.Cwd
		}
		ret, err := cmd.Output()
		if err != nil {
			ctx.Must(ctx.Error("Error executing command:", query.String(), err.Error(), string(ret)))
		}
		return string(ret)
	})
}
