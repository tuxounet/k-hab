package runtime

import "github.com/tuxounet/k-hab/utils"

func (l *RuntimeController) withIncusCmd(args ...string) (*utils.CmdCall, error) {

	return utils.WithCmdCall(l.ctx, "hab.incus.command.prefix", "hab.incus.command.name", args...)

}
