package definitions

import (
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
	default:
		return HabBaseDefinition{}, errors.New("base not found with name: " + name)
	}
}
