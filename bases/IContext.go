package bases

type IContext interface {
	Getwd() string
	GetConfigValue(key string) string
	GetCurrentConfig() map[string]string
	// GetHabConfig() HabConfig
	// SetHabConfig(HabConfig)
	GetSetupContainers() []SetupContainer
	GetLogger() ILogger
	GetSubLogger(name string, parent ILogger) ILogger
	GetController(key HabControllers) (IController, error)
}
