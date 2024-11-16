package plateform

import (
	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/controllers/dependencies"
	"github.com/tuxounet/k-hab/utils"
)

func (l *PlateformController) withIncusCmd(args ...string) (*utils.CmdCall, error) {

	return utils.WithCmdCall(l.ctx, "hab.incus.command.prefix", "hab.incus.command.name", args...)

}

func (l *PlateformController) withSystemCtlCmd(args ...string) (*utils.CmdCall, error) {

	return utils.WithCmdCall(l.ctx, "hab.commands.systemctl.prefix", "hab.commands.systemctl", args...)

}

func (l *PlateformController) withPsCmd(args ...string) (*utils.CmdCall, error) {

	return utils.WithCmdCall(l.ctx, "hab.commands.ps.prefix", "hab.commands.ps", args...)

}

func (l *PlateformController) getDependenciesController() (*dependencies.DependenciesController, error) {
	controller, err := l.ctx.GetController(bases.DependenciesController)
	if err != nil {
		return nil, err
	}
	dependencyController := controller.(*dependencies.DependenciesController)
	return dependencyController, nil
}
