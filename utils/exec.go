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

func WithCmdCall(anyMap any, cmdPrefixKey string, cmdNameKey string, args ...string) (*CmdCall, error) {

	cmd_prefix := GetMapValue(anyMap, cmdPrefixKey).(string)
	cmd := GetMapValue(anyMap, cmdNameKey).(string)

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
	}, nil

}

func OsExec(query *CmdCall) error {

	code, err := OsExecWithExitCode(query)
	if err != nil {
		return err
	}
	if code != 0 {
		return errors.New("command failed with exit code " + fmt.Sprint(code))
	}
	return nil

}

func OsExecWithExitCode(query *CmdCall) (int, error) {

	cmd := exec.Command(query.Command, query.Args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if query.Cwd != nil {
		cmd.Dir = *query.Cwd
	}

	err := cmd.Start()
	if err != nil {
		return -1, err
	}

	cmd.Wait()
	return cmd.ProcessState.ExitCode(), nil

}

func JsonCommandOutput[R any](query *CmdCall) (R, error) {
	var result R
	out, err := RawCommandOutput(query)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal([]byte(out), &result)
	if err != nil {
		return result, err
	}

	return result, nil

}

func RawCommandOutput(query *CmdCall) (string, error) {

	cmd := exec.Command(query.Command, query.Args...)
	if query.Cwd != nil {
		cmd.Dir = *query.Cwd
	}
	ret, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(ret), nil

}
