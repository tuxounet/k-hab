package containers

import (
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/controllers/dependencies"

	"github.com/tuxounet/k-hab/utils"
)

type ContainerModel struct {
	ctx  bases.IContext
	name string
	arch string

	ContainerConfig bases.HabContainerConfig
}

func NewContainerModel(name string, ctx bases.IContext, containerConfig bases.HabContainerConfig) *ContainerModel {

	return &ContainerModel{
		name:            name,
		ctx:             ctx,
		arch:            runtime.GOARCH,
		ContainerConfig: containerConfig,
	}
}

func (l *ContainerModel) withLxcCmd(args ...string) (*utils.CmdCall, error) {

	return utils.WithCmdCall(l.ctx.GetHabConfig(), "lxd.lxc.command.prefix", "lxd.lxc.command.name", args...)

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
		if container["name"] == l.name {
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
		if container["name"] == l.name {
			return container["status"].(string), nil
		}
	}
	return "unknown", nil

}

func (l *ContainerModel) Provision() error {

	containerExists, err := l.Present()
	if err != nil {
		return err
	}
	if !containerExists {
		conf := l.ContainerConfig.ToMap()

		containerImage := utils.GetMapValue(conf, "image").(string)

		lxcProfile := utils.GetMapValue(l.ctx.GetHabConfig(), "lxd.lxc.profile").(string)
		lxdCmd, err := l.withLxcCmd("init", containerImage, l.name, "--profile", lxcProfile)
		if err != nil {
			return err
		}

		cloudInit := utils.GetMapValue(conf, "cloud-init").(string)
		networkConfig := utils.GetMapValue(conf, "network-config").(string)

		if cloudInit != "" {
			sCloudInit, err := utils.UnTemplate(cloudInit, map[string]interface{}{
				"hab":       l.ctx.GetHabConfig(),
				"container": conf,
			})
			if err != nil {
				return err
			}
			userDataInclude := fmt.Sprintf(`--config=user.user-data=%s`, sCloudInit)
			lxdCmd.Args = append(lxdCmd.Args, userDataInclude)
		}

		if networkConfig != "" {
			sNetworkConfig, err := utils.UnTemplate(networkConfig, map[string]interface{}{
				"hab":       l.ctx.GetHabConfig(),
				"container": conf,
			})
			if err != nil {
				return err
			}
			userDataInclude := fmt.Sprintf(`--config=user.network-config=%s`, sNetworkConfig)
			lxdCmd.Args = append(lxdCmd.Args, userDataInclude)
		}

		err = utils.OsExec(lxdCmd)
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
		cmd, err := l.withLxcCmd("start", l.name)
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
			cmd, err := l.withLxcCmd("exec", l.name, "--", "cloud-init", "status", "--wait")
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

	cmd, err := l.withLxcCmd("exec", l.name, "--")
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

func (l *ContainerModel) Stop() error {

	status, err := l.Status()
	if err != nil {
		return err
	}
	if status == "Running" {
		cmd, err := l.withLxcCmd("stop", l.name)
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

func (l *ContainerModel) Unprovision() error {
	controller, err := l.ctx.GetController("DependenciesController")
	if err != nil {
		return err
	}
	dependencyController := controller.(*dependencies.DependenciesController)

	config := l.ctx.GetHabConfig()

	snapName := utils.GetMapValue(config, "lxd.snap").(string)
	present, err := dependencyController.InstalledSnap(snapName)
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
			cmd, err := l.withLxcCmd("delete", l.name)
			if err != nil {
				return err
			}
			err = utils.OsExec(cmd)
			if err != nil {
				return err
			}
		}
	}
	return nil

}
