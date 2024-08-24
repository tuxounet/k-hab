package bases

type ILogger interface {
	GetName() string
	CreateSubLogger(name string, parent ILogger) ILogger
	TraceF(format string, args ...interface{})
	DebugF(format string, args ...interface{})
	InfoF(format string, args ...interface{})
	WarnF(format string, args ...interface{})
	ErrorF(format string, args ...interface{})
}
