package bases

type IController interface {
	Install() error
	Uninstall() error
	Provision() error
	Start() error
	Deploy() error
	Undeploy() error
	Stop() error
	Rm() error
	Unprovision() error
	Nuke() error
}
