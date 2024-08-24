package bases

type IController interface {
	Provision() error
	Start() error
	Stop() error
	Rm() error
	Unprovision() error
	Nuke() error
}
