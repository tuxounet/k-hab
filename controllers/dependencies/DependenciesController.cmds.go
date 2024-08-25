package dependencies

import "github.com/tuxounet/k-hab/utils"

func (h *DependenciesController) withSnapCmd(args ...string) (*utils.CmdCall, error) {

	return utils.WithCmdCall(h.ctx, "hab.snap.command.prefix", "hab.snap.command.name", args...)

}
