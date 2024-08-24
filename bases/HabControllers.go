package bases

type HabControllers string

const (
	IngressController      HabControllers = "IngressController"
	EgressController       HabControllers = "EgressController"
	ContainersController   HabControllers = "ContainersController"
	BuilderController      HabControllers = "BuilderController"
	DependenciesController HabControllers = "DependenciesController"
	RuntimeController      HabControllers = "RuntimeController"
	ImagesController       HabControllers = "ImagesController"
)

func HabControllersLoadOrder() []HabControllers {
	return []HabControllers{
		DependenciesController,
		RuntimeController,
		BuilderController,
		ImagesController,
		ContainersController,
		EgressController,
		IngressController,
	}
}

func HabControllersUnloadOrder() []HabControllers {
	return []HabControllers{
		IngressController,
		EgressController,
		ContainersController,
		ImagesController,
		BuilderController,
		RuntimeController,
		DependenciesController,
	}

}
