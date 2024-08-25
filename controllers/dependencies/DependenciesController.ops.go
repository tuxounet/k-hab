package dependencies

import (
	"strings"

	"github.com/tuxounet/k-hab/utils"
)

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
	h.log.TraceF("Installing snap %s", name)
	cmd, err := h.withSnapCmd("install", name, mode)
	if err != nil {
		return err
	}

	err = utils.OsExec(cmd)
	if err != nil {
		return err
	}
	h.log.DebugF("Installed snap %s", name)
	return nil
}

func (h *DependenciesController) RemoveSnap(name string) error {

	installed, err := h.InstalledSnap(name)
	if err != nil {
		return err
	}
	if installed {

		h.log.TraceF("Removing snap %s", name)

		cmd, err := h.withSnapCmd("remove", name)
		if err != nil {
			return err
		}

		err = utils.OsExec(cmd)
		if err != nil {
			return err
		}
		h.log.DebugF("Removed snap %s", name)
	}
	return nil
}

func (h *DependenciesController) TakeSnapSnapshots(name string) error {
	h.log.TraceF("Taking snapshots for snap %s", name)
	cmd, err := h.withSnapCmd("save", name)
	if err != nil {
		return err
	}

	err = utils.OsExec(cmd)
	if err != nil {
		return err
	}

	h.log.DebugF("Snapshots taken for snap %s", name)
	return nil
}

func (h *DependenciesController) RemoveSnapSnapshots(name string) error {

	h.log.TraceF("Removing snapshots for snap %s", name)
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
	h.log.DebugF("Snapshots removed for snap %s", name)
	return nil
}

func (h *DependenciesController) ListSnapshots(name string) ([]string, error) {
	h.log.TraceF("Listing snapshots for snap %s", name)

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
		snap := strings.Fields(line)[1]
		if snap == name {

			snapshots = append(snapshots, id)
		}

	}

	h.log.DebugF("Snapshots for snap %s: %v", name, len(snapshots))
	return snapshots, nil

}

func (h *DependenciesController) ForgetSnapshot(name string, id string) error {
	h.log.TraceF("Forgetting snapshot %s for snap %s", id, name)

	cmd, err := h.withSnapCmd("forget", id, name)
	if err != nil {
		return err
	}

	err = utils.OsExec(cmd)
	if err != nil {
		return err
	}
	h.log.DebugF("Forgotten snapshot %s for snap %s", id, name)
	return nil

}
