package utils

import (
	"testing"
)

func TestTTContext(t *testing.T) {
	ctx := NewScopeContext(true, "TESTING")
	if ctx == nil {
		t.Fatalf("Expected context, got nil")
	}
}

func TestTTContextError(t *testing.T) {
	ctx := NewTestContext()
	err := ctx.Error("somtething went wrong")
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestTTMustRecoverPanic(t *testing.T) {
	ctx := NewTestContext()

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Expected panic, got nil")
		}
	}()
	ctx.Must(ctx.Error("something went wrong"))
}

func TestTTScope(t *testing.T) {

	called := false
	ctx := NewTestContext()
	err := ctx.Scope("test", "scope", func(ctx *ScopeContext) {
		called = true
	})
	if err != nil {
		t.Fatalf("Expected nil, got %v", err)
	}
	if !called {
		t.Fatalf("Expected true, got false")
	}
}

func TestTTScopeReturn(t *testing.T) {

	ctx := NewTestContext()
	ret := ScopingWithReturn(ctx, "test", "scope", func(ctx *ScopeContext) string {
		return "test"
	})

	if ret != "test" {
		t.Fatalf("Expected 'test', got '%s'", ret)
	}
}
func TestTTSubScope(t *testing.T) {

	ctx := NewTestContext()
	ret := ScopingWithReturn(ctx, "test", "scope", func(ctx *ScopeContext) string {
		ret := ""
		ctx.Scope("sub", "scope", func(ctx *ScopeContext) {
			ret = "test"
		})

		return ret
	})

	if ret != "test" {
		t.Fatalf("Expected 'test', got '%s'", ret)
	}
}

func TestTTSubScopeReturn(t *testing.T) {

	ctx := NewTestContext()
	ret := ScopingWithReturn(ctx, "test", "scope", func(ctx *ScopeContext) string {

		ret := ScopingWithReturn(ctx, "sub", "scope", func(ctx *ScopeContext) string {
			exp := "TESTING.sub/scope"
			if ctx.Name != exp {
				t.Fatalf("Expected '%s', got '%s'", exp, ctx.Name)
			}

			return "test"
		})

		return ret
	})

	if ret != "test" {
		t.Fatalf("Expected 'test', got '%s'", ret)
	}
}
