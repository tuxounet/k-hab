package logger

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestTTLogLevels(t *testing.T) {
	// Test all log levels
	logger := NewLogger(context.TODO(), "test")

	name := logger.GetName()
	if name != "test" {
		t.Fatalf("Expected 'test', got '%s'", name)
	}

	logger.SetLevel(logrus.WarnLevel)

	logger.TraceF("trace")
	logger.DebugF("debug")
	logger.InfoF("info")
	logger.WarnF("warn")
	logger.ErrorF("error")
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Expected panic")
		}
	}()
	logger.PanicF("panic")

	//recover panic

}

func TestTTSubLogger(t *testing.T) {
	// Test sublogger
	logger := NewLogger(context.TODO(), "test")
	sublogger := logger.CreateSubLogger("sub", logger)

	name := sublogger.GetName()
	if name != "test.sub" {
		t.Fatalf("Expected 'test.sub', got '%s'", name)
	}

}
