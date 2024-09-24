package bases

type BaseController struct {
	Ctx *IContext
}

func (b *BaseController) Provision() error {
	return nil
}
func (b *BaseController) Start() error {
	return nil
}
func (b *BaseController) Deploy() error {
	return nil
}
func (b *BaseController) Undeploy() error {
	return nil
}
func (b *BaseController) Stop() error {
	return nil
}
func (b *BaseController) Rm() error {
	return nil
}
func (b *BaseController) Unprovision() error {
	return nil
}
func (b *BaseController) Nuke() error {
	return nil
}
