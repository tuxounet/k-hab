package containers

import (
	"errors"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/tuxounet/k-hab/bases"

	"github.com/tuxounet/k-hab/utils"
)

type ContainerModel struct {
	ctx            bases.IContext
	containersPath string
	Name           string
	Arch           string

	ContainerConfig bases.SetupContainer
}

func NewContainerModel(name string, ctx bases.IContext, containerConfig bases.SetupContainer, containersPath string) *ContainerModel {

	return &ContainerModel{
		ctx:             ctx,
		containersPath:  containersPath,
		Name:            name,
		Arch:            runtime.GOARCH,
		ContainerConfig: containerConfig,
	}
}

func (l *ContainerModel) Present() (bool, error) {

	cmd, err := l.withLxcCmd("list", "--format", "json")
	if err != nil {
		return false, err
	}
	out, err := utils.JsonCommandOutput[[]map[string]interface{}](cmd)
	if err != nil {
		return false, err
	}

	for _, container := range out {
		if container["name"] == l.Name {
			return true, nil
		}
	}
	return false, nil

}

func (l *ContainerModel) Status() (string, error) {
	cmd, err := l.withLxcCmd("list", "--format", "json")
	if err != nil {
		return "", err
	}

	out, err := utils.JsonCommandOutput[[]map[string]interface{}](cmd)

	if err != nil {
		return "", err
	}
	for _, container := range out {
		if container["name"] == l.Name {
			return container["status"].(string), nil
		}
	}
	return "unknown", nil

}

func (l *ContainerModel) Provision() error {
	imagesController, err := l.getImagesController()
	if err != nil {
		return err
	}

	baseChanged, err := imagesController.EnsureImage(l.ContainerConfig.Base)
	if err != nil {
		return err
	}

	pfCmd, err := l.getLaunchCmd()
	if err != nil {
		return err
	}

	containerFile := path.Join(l.containersPath, l.Name)

	_, err = os.Stat(containerFile)
	if err == nil {
		body, err := os.ReadFile(containerFile)
		if err != nil {
			return err
		}
		if pfCmd.String() != string(body) {
			//Change
			baseChanged = true
		}

	}

	containerExists, err := l.Present()
	if err != nil {
		return err
	}
	if baseChanged {
		err = l.Stop()
		if err != nil {
			return err
		}
		err = l.Unprovision()
		if err != nil {
			return err
		}
		containerExists = false
	}

	if !containerExists {

		err = utils.OsExec(pfCmd)
		if err != nil {
			return err
		}
		//write commandline to file
		err = os.WriteFile(containerFile, []byte(pfCmd.String()), 0644)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (l *ContainerModel) Start() error {

	status, err := l.Status()
	if err != nil {
		return err
	}

	if status != "Running" {
		cmd, err := l.withLxcCmd("start", l.Name)
		if err != nil {
			return err
		}
		err = utils.OsExec(cmd)
		if err != nil {
			return err
		}
		return nil

	}
	return nil

}

func (l *ContainerModel) Deploy() error {

	err := l.Start()
	if err != nil {
		return err
	}
	err = l.Exec("/etc/deploy.sh")
	if err != nil {
		return err
	}

	return nil

}

func (l *ContainerModel) WaitReady() error {

	timeout := 30 * time.Second

	// Heure de fin
	heureFin := time.Now().Add(timeout)

	// Boucle jusqu'à l'heure de fin
	for time.Now().Before(heureFin) {

		status, err := l.Status()
		if err != nil {
			return err
		}
		if status == "Running" {
			cmd, err := l.withLxcCmd("exec", l.Name, "--", "cloud-init", "status", "--wait")
			if err != nil {
				return err
			}
			code, err := utils.OsExecWithExitCode(cmd)
			if err != nil {
				return err
			}
			if code == 2 || code == 0 {
				return nil
			}
		}
		time.Sleep(1 * time.Second)
	}
	return errors.New("timeout to waiting ready")

}

func (l *ContainerModel) Exec(command ...string) error {

	cmd, err := l.withLxcCmd("exec", l.Name, "--")
	if err != nil {
		return err
	}
	cmd.Args = append(cmd.Args, command...)
	err = utils.OsExec(cmd)
	if err != nil {
		return err
	}
	return nil

}

func (l *ContainerModel) Shell() error {
	shellCmd := l.ContainerConfig.Shell
	err := l.Exec(shellCmd)
	if err != nil {
		l.ctx.GetLogger().ErrorF("Error while executing shell command: %s", err)
	}
	return nil
}

func (l *ContainerModel) Stop() error {

	status, err := l.Status()
	if err != nil {
		return err
	}
	if status == "Running" {
		cmd, err := l.withLxcCmd("stop", l.Name)
		if err != nil {
			return err
		}
		err = utils.OsExec(cmd)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *ContainerModel) Undeploy() error {

	err := l.Start()
	if err != nil {
		return err
	}
	err = l.Exec("/etc/undeploy.sh")
	if err != nil {
		return err
	}

	return nil
}

func (l *ContainerModel) Unprovision() error {

	plateformController, err := l.getPlateformController()
	if err != nil {
		return err
	}
	present, err := plateformController.IsPresent()
	if err != nil {
		return err
	}

	if present {

		containerExists, err := l.Present()
		if err != nil {
			return err
		}

		if containerExists {
			err = l.Stop()
			if err != nil {
				return err
			}
			cmd, err := l.withLxcCmd("delete", l.Name)
			if err != nil {
				return err
			}
			err = utils.OsExec(cmd)
			if err != nil {
				return err
			}

		}
	}
	containerFile := path.Join(l.containersPath, l.Name)
	if _, err := os.Stat(containerFile); err == nil {
		err = os.Remove(containerFile)
		if err != nil {
			return err
		}
	}

	return nil

}
