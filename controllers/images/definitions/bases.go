package definitions

import (
	_ "embed"
	"errors"
)

type HabBaseDefinition struct {
	Name          string `yaml:"name"`
	Builder       string `yaml:"builder"`
	CloudInit     string `yaml:"cloudinit"`
	NetworkConfig string `yaml:"networkconfig"`
}

func GetImageBase(name string) (HabBaseDefinition, error) {
	switch name {
	case "alpine":
		return AlpineImageBase, nil
	case "ubuntu":
		return UbuntuImageBase, nil
	default:
		return HabBaseDefinition{}, errors.New("base not found with name: " + name)
	}
}

//go:embed alpine/distrobuilder.yml
var alpineBuilderDefinition string

//go:embed alpine/cloud-init.yml
var alpineCloudInit string

//go:embed alpine/network-config.yml
var alpineNetworkConfig string

var AlpineImageBase = HabBaseDefinition{
	Name:          "alpine",
	Builder:       alpineBuilderDefinition,
	CloudInit:     alpineCloudInit,
	NetworkConfig: alpineNetworkConfig,
}

//go:embed ubuntu/distrobuilder.yml
var ubuntuBuilderDefinition string

//go:embed ubuntu/cloud-init.yml
var ubuntuCloudInit string

//go:embed ubuntu/network-config.yml
var ubuntuNetworkConfig string

var UbuntuImageBase = HabBaseDefinition{
	Name:          "ubuntu",
	Builder:       ubuntuBuilderDefinition,
	CloudInit:     ubuntuCloudInit,
	NetworkConfig: ubuntuNetworkConfig,
}
