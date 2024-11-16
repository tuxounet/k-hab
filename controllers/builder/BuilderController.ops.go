package builder

import (
	"os"
	"path"

	"github.com/tuxounet/k-hab/utils"
)

const (
	imageMetadataPackage = "incus.tar.xz"
	imageRootfsPackage   = "rootfs.squashfs"
)

type DistroBuilderResult struct {
	Built           bool
	MetadataPackage string
	RootfsPackage   string
}

func (l *BuilderController) BuildDistro(name string, builderConfig string) (*DistroBuilderResult, error) {

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

		distroBuildFile := path.Join(distroFolder, "distro.yaml")
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

	return &DistroBuilderResult{
		MetadataPackage: metadataPackage,
		RootfsPackage:   rootfsPackage,
		Built:           built,
	}, nil

}

func (l *BuilderController) RemoveCache(name string) error {
	builderPath, err := l.getImageBuildPath()
	if err != nil {
		return err
	}
	distroFolder := path.Join(builderPath, name)

	cmd, err := utils.WithCmdCall(l.ctx, "hab.rm.prefix", "hab.rm.name", "-rf", distroFolder)
	if err != nil {
		return err
	}

	return utils.OsExec(cmd)

}

func (l *BuilderController) ConfigHasChanged(name string, expectedConfig string) (bool, error) {
	builderPath, err := l.getImageBuildPath()
	if err != nil {
		return false, err
	}
	distroFolder := path.Join(builderPath, name)
	distroBuildFile := path.Join(distroFolder, "distro.yaml")

	if _, err := os.Stat(distroBuildFile); os.IsNotExist(err) {
		return true, nil
	}
	metadataPackage := path.Join(distroFolder, imageMetadataPackage)
	rootfsPackage := path.Join(distroFolder, imageRootfsPackage)

	if _, err := os.Stat(metadataPackage); os.IsNotExist(err) {
		return true, nil
	}

	if _, err := os.Stat(rootfsPackage); os.IsNotExist(err) {
		return true, nil
	}

	currentConfig, err := os.ReadFile(distroBuildFile)
	if err != nil {
		return false, err
	}
	config := string(currentConfig)

	return config != expectedConfig, nil

}
