package main

import (
	"testing"

	"github.com/tuxounet/k-hab/context"
)

func TestBuildCommand(t *testing.T) {
	cmd := buildCommand("aaa", "bbb", context.DeployVerb)
	if cmd.Name != "aaa" {
		t.Errorf("cmd.Name = %s; want aaa", cmd.Name)
	}
	if cmd.Usage != "bbb" {
		t.Errorf("cmd.Description = %s; want bbb", cmd.Description)
	}
	if cmd.Action == nil {
		t.Errorf("cmd.Action = nil; want not nil")
	}

}
