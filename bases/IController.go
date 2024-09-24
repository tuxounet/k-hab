package bases

type IController interface {
	Provision() error
	Start() error
	Deploy() error
	Undeploy() error
	Stop() error
	Rm() error
	Unprovision() error
	Nuke() error
}
