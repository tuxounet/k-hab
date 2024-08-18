package host

import (
	"strings"

	"github.com/tuxounet/k-hab/utils"
)

type SnapPackages struct {
	scopeBase string
	habConfig map[string]interface{}
}

func NewSnapPackages(config interface{}) *SnapPackages {
	return &SnapPackages{
		scopeBase: "SnapPackages",
		habConfig: config.(map[string]interface{}),
	}
}

type snapCmd struct {
	cmd  string
	args []string
}

func (h *SnapPackages) withSnapCmd(ctx *utils.ScopeContext, args ...string) *utils.CmdCall {

	return utils.ScopingWithReturn(ctx, h.scopeBase, "Present", func(ctx *utils.ScopeContext) *utils.CmdCall {
		return utils.WithCmdCallBuilder(ctx, h.habConfig, "snap.command.prefix", "snap.command.name", args...)

	})
}

func (h *SnapPackages) InstalledSnap(ctx *utils.ScopeContext, name string) bool {
	return utils.ScopingWithReturn(ctx, h.scopeBase, "InstalledSnap", func(ctx *utils.ScopeContext) bool {
		out := utils.CommandSyncOutput(ctx, h.withSnapCmd(ctx, "list"))

		//parse out to array of strings for each line
		lines := strings.Split(strings.TrimSpace(out), "\n")

		for i := 1; i < len(lines); i++ {
			line := lines[i]
			if strings.HasPrefix(line, name+" ") {
				return true
			}
		}
		return false

	})

}

func (h *SnapPackages) InstallSnap(ctx *utils.ScopeContext, name string, mode string) error {
	return ctx.Scope(h.scopeBase, "InstallSnap", func(ctx *utils.ScopeContext) {
		ctx.Must(utils.ExecSyncOutput(ctx, h.withSnapCmd(ctx, "install", name, mode)))
	})

}

func (h *SnapPackages) RemoveSnap(ctx *utils.ScopeContext, name string) error {
	return ctx.Scope(h.scopeBase, "RemoveSnap", func(ctx *utils.ScopeContext) {
		ctx.Must(utils.ExecSyncOutput(ctx, h.withSnapCmd(ctx, "remove", name)))
	})

}

func (h *SnapPackages) RemoveSnapSnapshots(ctx *utils.ScopeContext, name string) error {
	return ctx.Scope(h.scopeBase, "RemoveSnapSnapshots", func(ctx *utils.ScopeContext) {
		snapshots := h.ListSnapshots(ctx, name)

		for _, snap := range snapshots {
			err := h.ForgetSnapshot(ctx, name, snap)
			ctx.Must(err)
		}
	})
}

func (h *SnapPackages) ListSnapshots(ctx *utils.ScopeContext, name string) []string {
	return utils.ScopingWithReturn(ctx, h.scopeBase, "ListSnapshots", func(ctx *utils.ScopeContext) []string {
		out := utils.CommandSyncOutput(ctx, h.withSnapCmd(ctx, "saved"))

		if !strings.HasPrefix(out, "Set") {
			return make([]string, 0)
		}

		snapshots := make([]string, 0)
		//parse out to array of strings for each line
		lines := strings.Split(strings.TrimSpace(out), "\n")

		for i := 1; i < len(lines); i++ {
			line := lines[i]
			id := strings.Fields(line)[0]
			snapshots = append(snapshots, id)

		}

		return snapshots
	})
}

func (h *SnapPackages) ForgetSnapshot(ctx *utils.ScopeContext, name string, id string) error {
	return ctx.Scope(h.scopeBase, "ForgetSnapshot", func(ctx *utils.ScopeContext) {
		ctx.Must(utils.ExecSyncOutput(ctx, h.withSnapCmd(ctx, "forget", id, name)))
	})
}
