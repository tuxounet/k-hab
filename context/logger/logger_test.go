package logger

import (
	"context"
	"testing"
)

func TestTTLogLevels(t *testing.T) {
	// Test all log levels
	logger := NewLogger(context.TODO())
	logger.TraceF("trace")
	logger.DebugF("debug")
	logger.InfoF("info")
	logger.WarnF("warn")
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Expected panic")
		}
	}()
	logger.PanicF("panic")

	//recover panic

}
