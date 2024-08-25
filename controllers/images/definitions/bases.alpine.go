package definitions

import (
	_ "embed"
)

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
