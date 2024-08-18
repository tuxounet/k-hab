package utils

import "testing"

func TestTTLogLevels(t *testing.T) {
	ctx := NewTestContext()
	ctx.Log.DebugF("debug")
	ctx.Log.InfoF("info")
	ctx.Log.WarnF("warn")

}
