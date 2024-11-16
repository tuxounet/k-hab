package containers

import (
	"fmt"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/controllers/images"

	"github.com/tuxounet/k-hab/utils"
)

func (l *ContainerModel) withIncusCmd(args ...string) (*utils.CmdCall, error) {

	return utils.WithCmdCall(l.ctx, "hab.incus.command.prefix", "hab.incus.command.name", args...)

}

func (l *ContainerModel) getLaunchCmd() (*utils.CmdCall, error) {
	imgcontroller, err := l.ctx.GetController(bases.ImagesController)
	if err != nil {
		return nil, err
	}
	imagesController := imgcontroller.(*images.ImagesController)
	image, err := imagesController.GetImage(l.ContainerConfig.Base)
	if err != nil {
		return nil, err
	}
	incusProfile := l.ctx.GetConfigValue("hab.incus.profile")
	incusCmd, err := l.withIncusCmd("init", l.ContainerConfig.Base, l.Name, "--profile", incusProfile)
	if err != nil {
		return nil, err
	}

	if image.Definition.CloudInit != "" {
		sCloudInit, err := utils.UnTemplate(image.Definition.CloudInit, map[string]interface{}{
			"config":    l.ctx.GetCurrentConfig(),
			"container": l.ContainerConfig.ToMap(),
		})
		if err != nil {
			return nil, err
		}
		userDataInclude := fmt.Sprintf(`--config=user.user-data=%s`, sCloudInit)
		incusCmd.Args = append(incusCmd.Args, userDataInclude)
	}

	if image.Definition.NetworkConfig != "" {
		sNetworkConfig, err := utils.UnTemplate(image.Definition.NetworkConfig, map[string]interface{}{
			"config":    l.ctx.GetCurrentConfig(),
			"container": l.ContainerConfig.ToMap(),
		})
		if err != nil {
			return nil, err
		}
		userDataInclude := fmt.Sprintf(`--config=user.network-config=%s`, sNetworkConfig)
		incusCmd.Args = append(incusCmd.Args, userDataInclude)
	}
	return incusCmd, nil
}
