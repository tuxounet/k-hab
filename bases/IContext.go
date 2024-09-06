package bases

type IContext interface {
	GetConfigValue(key string) string
	SetConfigValue(key string, value string)
	GetStorageRoot() (string, error)
	GetCurrentConfig() map[string]string
	GetSetupContainers() []SetupContainer
	GetLogger() ILogger
	GetSubLogger(name string, parent ILogger) ILogger
	GetController(key HabControllers) (IController, error)
}
