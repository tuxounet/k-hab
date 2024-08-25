package bases

type IContext interface {
	Getwd() string
	GetHabConfig() HabConfig
	SetHabConfig(HabConfig)
	GetContainersConfig() []HabContainerConfig
	GetLogger() ILogger
	GetSubLogger(name string, parent ILogger) ILogger
	GetController(key HabControllers) (IController, error)
}
