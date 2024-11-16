package dependencies

import "github.com/tuxounet/k-hab/utils"

func (h *DependenciesController) withSnapCmd(args ...string) (*utils.CmdCall, error) {

	return utils.WithCmdCall(h.ctx, "hab.commands.snap.prefix", "hab.commands.snap", args...)

}

func (h *DependenciesController) withAptCmd(args ...string) (*utils.CmdCall, error) {

	return utils.WithCmdCall(h.ctx, "hab.apt.command.prefix", "hab.apt.command.name", args...)

}
func (h *DependenciesController) withDpkgCommand(args ...string) (*utils.CmdCall, error) {

	return utils.WithCmdCall(h.ctx, "hab.commands.dpkg.prefix", "hab.commands.dpkg", args...)

}
