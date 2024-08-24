package bases

type HabControllers string

const (
	IngressController      HabControllers = "IngressController"
	ContainersController   HabControllers = "ContainersController"
	BuilderController      HabControllers = "BuilderController"
	DependenciesController HabControllers = "DependenciesController"
)
