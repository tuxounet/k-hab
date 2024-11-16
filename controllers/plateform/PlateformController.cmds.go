package plateform

import (
	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/controllers/dependencies"
	"github.com/tuxounet/k-hab/utils"
)

func (l *PlateformController) withLxcCmd(args ...string) (*utils.CmdCall, error) {

	return utils.WithCmdCall(l.ctx, "hab.plateform.command.prefix", "hab.plateform.command", args...)

}

func (l *PlateformController) getDependenciesController() (*dependencies.DependenciesController, error) {
	controller, err := l.ctx.GetController(bases.DependenciesController)
	if err != nil {
		return nil, err
	}
	dependencyController := controller.(*dependencies.DependenciesController)
	return dependencyController, nil
}
