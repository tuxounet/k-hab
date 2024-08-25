package runtime

import "github.com/tuxounet/k-hab/utils"

func (l *RuntimeController) withLxdCmd(args ...string) (*utils.CmdCall, error) {

	return utils.WithCmdCall(l.ctx, "hab.lxd.command.prefix", "hab.lxd.command.name", args...)

}

func (l *RuntimeController) withLxcCmd(args ...string) (*utils.CmdCall, error) {

	return utils.WithCmdCall(l.ctx, "hab.lxd.lxc.command.prefix", "hab.lxd.lxc.command.name", args...)

}
