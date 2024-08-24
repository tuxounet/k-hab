package runtime

import "github.com/tuxounet/k-hab/utils"

func (l *RuntimeController) withLxdCmd(args ...string) (*utils.CmdCall, error) {

	return utils.WithCmdCall(l.ctx.GetHabConfig(), "lxd.command.prefix", "lxd.command.name", args...)

}

func (l *RuntimeController) withLxcCmd(args ...string) (*utils.CmdCall, error) {

	return utils.WithCmdCall(l.ctx.GetHabConfig(), "lxd.lxc.command.prefix", "lxd.lxc.command.name", args...)

}
