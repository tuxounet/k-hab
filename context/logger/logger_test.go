package logger_test

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/tuxounet/k-hab/context/logger"
)

func TestTTLogLevels(t *testing.T) {
	// Test all log levels
	log := logger.NewLogger(context.TODO(), "test")

	name := log.GetName()
	if name != "test" {
		t.Fatalf("Expected 'test', got '%s'", name)
	}

	log.SetLevel(logrus.WarnLevel)

	log.TraceF("trace")
	log.DebugF("debug")
	log.InfoF("info")
	log.WarnF("warn")
	log.ErrorF("error")
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Expected panic")
		}
	}()
	log.PanicF("panic")

	//recover panic

}

func TestTTSubLogger(t *testing.T) {
	// Test sublogger
	log := logger.NewLogger(context.TODO(), "test")
	sublogger := log.CreateSubLogger("sub", log)

	name := sublogger.GetName()
	if name != "test.sub" {
		t.Fatalf("Expected 'test.sub', got '%s'", name)
	}

}
