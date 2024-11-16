package containers

import (
	"fmt"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/controllers/images"
	"github.com/tuxounet/k-hab/controllers/plateform"

	"github.com/tuxounet/k-hab/utils"
)

func (l *ContainerModel) withLxcCmd(args ...string) (*utils.CmdCall, error) {

	return utils.WithCmdCall(l.ctx, "hab.plateform.command.prefix", "hab.plateform.command", args...)

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
	pfProfile := l.ctx.GetConfigValue("hab.plateform.profile")
	pfCmd, err := l.withLxcCmd("init", l.ContainerConfig.Base, l.Name, "--profile", pfProfile)
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
		pfCmd.Args = append(pfCmd.Args, userDataInclude)
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
		pfCmd.Args = append(pfCmd.Args, userDataInclude)
	}
	return pfCmd, nil
}

func (l *ContainerModel) getPlateformController() (*plateform.PlateformController, error) {
	controller, err := l.ctx.GetController(bases.PlateformController)
	if err != nil {
		return nil, err
	}
	plateformController := controller.(*plateform.PlateformController)
	return plateformController, nil
}
