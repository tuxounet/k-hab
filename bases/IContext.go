package bases

type IContext interface {
	Getwd() string
	GetConfigValue(key string) string
	SetConfigValue(key string, value string)
	GetCurrentConfig() map[string]string
	GetSetupContainers() []SetupContainer
	GetLogger() ILogger
	GetSubLogger(name string, parent ILogger) ILogger
	GetController(key HabControllers) (IController, error)
}
