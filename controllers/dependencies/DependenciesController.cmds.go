package dependencies

import "github.com/tuxounet/k-hab/utils"

func (h *DependenciesController) withSnapCmd(args ...string) (*utils.CmdCall, error) {

	return utils.WithCmdCall(h.ctx, "hab.commands.snap.prefix", "hab.commands.snap", args...)

}
