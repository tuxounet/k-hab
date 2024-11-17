package builder

import (
	"os"
	"path"

	"github.com/tuxounet/k-hab/utils"
)

const (
	imageMetadataPackage = "incus.tar.xz"
	imageRootfsPackage   = "rootfs.squashfs"
	distroStateFile      = "distro.yaml"
)

type DistroBuilderResult struct {
	Built           bool
	MetadataPackage string
	RootfsPackage   string
}

func (l *BuilderController) BuildDistro(name string, builderConfig string) (*DistroBuilderResult, error) {
	l.log.TraceF("Build distro %s", name)

	builderPath, err := l.getImageBuildPath()
	if err != nil {
		return nil, err
	}
	distroFolder := path.Join(builderPath, name)
	os.MkdirAll(distroFolder, 0755)
	built := false
	changed, err := l.ConfigHasChanged(name, builderConfig)
	if err != nil {
		return nil, err
	}

	if changed {

		distroBuildFile := path.Join(distroFolder, distroStateFile)
		err = os.WriteFile(distroBuildFile, []byte(builderConfig), 0644)
		if err != nil {
			return nil, err
		}

		cmd, err := l.withDistroBuilderCmd("build-incus", distroBuildFile)
		if err != nil {
			return nil, err
		}
		cmd.Cwd = &distroFolder

		err = utils.OsExec(cmd)
		if err != nil {
			return nil, err
		}

	}

	metadataPackage := path.Join(distroFolder, imageMetadataPackage)
	rootfsPackage := path.Join(distroFolder, imageRootfsPackage)

	l.log.DebugF("disto built : %v", built)

	return &DistroBuilderResult{
		MetadataPackage: metadataPackage,
		RootfsPackage:   rootfsPackage,
		Built:           built,
	}, nil

}

func (l *BuilderController) RemoveCache(name string) error {
	l.log.TraceF("Remove cache for distro %s", name)
	builderPath, err := l.getImageBuildPath()
	if err != nil {
		return err
	}
	distroFolder := path.Join(builderPath, name)

	cmd, err := utils.WithCmdCall(l.ctx, "hab.commands.rm.prefix", "hab.commands.rm", "-rf", distroFolder)
	if err != nil {
		return err
	}

	err = utils.OsExec(cmd)
	if err != nil {
		return err
	}
	l.log.DebugF("Cache removed for distro %s", name)
	return nil
}

func (l *BuilderController) ConfigHasChanged(name string, expectedConfig string) (bool, error) {
	l.log.TraceF("Check if config has changed for distro %s", name)

	builderPath, err := l.getImageBuildPath()
	if err != nil {
		return false, err
	}
	distroFolder := path.Join(builderPath, name)
	distroBuildFile := path.Join(distroFolder, distroStateFile)

	if _, err := os.Stat(distroBuildFile); os.IsNotExist(err) {
		l.log.DebugF("No previous build for distro %s, changed", name)
		return true, nil
	}
	metadataPackage := path.Join(distroFolder, imageMetadataPackage)
	rootfsPackage := path.Join(distroFolder, imageRootfsPackage)

	if _, err := os.Stat(metadataPackage); os.IsNotExist(err) {
		l.log.DebugF("No previous metadata package for distro %s, changed", name)
		return true, nil
	}

	if _, err := os.Stat(rootfsPackage); os.IsNotExist(err) {
		l.log.DebugF("No previous rootfsPackage for distro %s, changed", name)
		return true, nil
	}

	currentConfig, err := os.ReadFile(distroBuildFile)
	if err != nil {
		return false, err
	}
	config := string(currentConfig)

	changed := config != expectedConfig
	l.log.DebugF("Config has changed for distro %s: %v", name, changed)
	return changed, nil

}
