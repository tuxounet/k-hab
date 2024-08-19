package utils

import (
	"context"
	"errors"
	"fmt"
)

type ScopingCallFunc func(ctx *ScopeContext)
type ScopingWithReturnCallFunc[R any] func(ctx *ScopeContext) (R, error)
type ScopingWithReturnNoErrorCallFunc[R any] func(ctx *ScopeContext) R

type ScopeContext struct {
	Name   string
	Log    *Logger
	parent *ScopeContext
}

func NewScopeContext(quiet bool, entrypoint string) *ScopeContext {

	ctx := context.Background()
	log := NewLogger(quiet, ctx)

	return &ScopeContext{
		Name:   entrypoint,
		Log:    log,
		parent: nil,
	}
}

func NewTestContext() *ScopeContext {

	ctx := context.Background()
	log := NewLogger(false, ctx)

	return &ScopeContext{
		Name:   "TESTING",
		Log:    log,
		parent: nil,
	}
}

func (s *ScopeContext) Must(err error) {
	if err != nil {
		s.Log.PanicF("🛑\tfailure: %v", err)
	}
}

func (s *ScopeContext) Error(args ...string) error {
	return errors.New(s.Name + " FAILURE: " + fmt.Sprintln(args))
}
func (s *ScopeContext) Scope(prefix string, name string, f ScopingCallFunc) error {
	scopeName := ""
	if s.parent != nil {
		scopeName = s.parent.Name + "." + prefix + "/" + name
	} else {
		scopeName = prefix + "/" + name
	}
	log := s.Log.CreateScopeLogger(scopeName, map[string]interface{}{})
	subScope := &ScopeContext{
		Name:   scopeName,
		Log:    log,
		parent: s,
	}

	log.TraceF("▶️")
	f(subScope)
	log.TraceF("◀️")
	return nil
}

func ScopingWithReturn[R any](s *ScopeContext, prefix string, name string, f ScopingWithReturnNoErrorCallFunc[R]) R {
	scopeName := ""
	if s.parent != nil {
		scopeName = s.parent.Name + "." + prefix + "/" + name
	} else {
		scopeName = prefix + "/" + name
	}
	log := s.Log.CreateScopeLogger(scopeName, map[string]interface{}{})
	subScope := &ScopeContext{
		Name:   scopeName,
		Log:    log,
		parent: s,
	}

	log.TraceF("▶️")
	out := f(subScope)

	log.TraceF("◀️")
	return out
}
