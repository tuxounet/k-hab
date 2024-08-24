package bases

type IContext interface {
	GetHabConfig() HabConfig
	SetHabConfig(HabConfig)
	GetController(key HabControllers) (IController, error)
}
