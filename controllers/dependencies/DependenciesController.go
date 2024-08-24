package dependencies

import (
	"strings"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/utils"
)

type DependenciesController struct {
	bases.BaseController
	ctx bases.IContext
}

func NewDependenciesController(ctx bases.IContext) *DependenciesController {
	return &DependenciesController{
		ctx: ctx,
	}
}

func (h *DependenciesController) withSnapCmd(args ...string) (*utils.CmdCall, error) {
	habConfig := h.ctx.GetHabConfig()
	return utils.WithCmdCall(habConfig, "snap.command.prefix", "snap.command.name", args...)

}

func (h *DependenciesController) InstalledSnap(name string) (bool, error) {

	cmd, err := h.withSnapCmd("list")
	if err != nil {
		return false, err
	}
	out, err := utils.RawCommandOutput(cmd)
	if err != nil {
		return false, err
	}

	//parse out to array of strings for each line
	lines := strings.Split(strings.TrimSpace(out), "\n")

	for i := 1; i < len(lines); i++ {
		line := lines[i]
		if strings.HasPrefix(line, name+" ") {
			return true, nil
		}
	}
	return false, nil

}

func (h *DependenciesController) InstallSnap(name string, mode string) error {
	cmd, err := h.withSnapCmd("install", name, mode)
	if err != nil {
		return err
	}

	err = utils.OsExec(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (h *DependenciesController) RemoveSnap(name string) error {
	cmd, err := h.withSnapCmd("remove", name)
	if err != nil {
		return err
	}

	err = utils.OsExec(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (h *DependenciesController) TakeSnapSnapshots(name string) error {
	cmd, err := h.withSnapCmd("save", name)
	if err != nil {
		return err
	}

	err = utils.OsExec(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (h *DependenciesController) RemoveSnapSnapshots(name string) error {

	snapshots, err := h.ListSnapshots(name)
	if err != nil {
		return err
	}

	for _, snap := range snapshots {
		err := h.ForgetSnapshot(name, snap)
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *DependenciesController) ListSnapshots(name string) ([]string, error) {
	cmd, err := h.withSnapCmd("saved")
	if err != nil {
		return nil, err
	}
	out, err := utils.RawCommandOutput(cmd)
	if err != nil {
		return nil, err
	}

	if !strings.HasPrefix(out, "Set") {
		return make([]string, 0), nil
	}

	snapshots := make([]string, 0)
	//parse out to array of strings for each line
	lines := strings.Split(strings.TrimSpace(out), "\n")

	for i := 1; i < len(lines); i++ {
		line := lines[i]
		id := strings.Fields(line)[0]
		snapshots = append(snapshots, id)

	}

	return snapshots, nil

}

func (h *DependenciesController) ForgetSnapshot(name string, id string) error {
	cmd, err := h.withSnapCmd("forget", id, name)
	if err != nil {
		return err
	}

	err = utils.OsExec(cmd)
	if err != nil {
		return err
	}
	return nil

}
