package config

import (
	_ "embed"
	"errors"

	"github.com/tuxounet/k-hab/bases"
	"github.com/tuxounet/k-hab/utils"
)

//go:embed templates/hab.yaml
var defaultHabConfig string

//go:embed templates/containers.yaml
var defaultContainersConfig string

//go:embed templates/images.yaml
var defaultImagesConfig string

type Config struct {
	HabConfig        bases.HabConfig
	ContainersConfig []bases.HabContainerConfig
	ImagesConfig     []bases.HabImageConfig
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() error {

	habConfig, err := utils.LoadYamlFromString[bases.HabConfig](defaultHabConfig)
	if err != nil {
		return err
	}
	containersConfig, err := utils.LoadYamlFromString[[]bases.HabContainerConfig](defaultContainersConfig)
	if err != nil {
		return err
	}

	imagesConfig, err := utils.LoadYamlFromString[[]bases.HabImageConfig](defaultImagesConfig)
	if err != nil {
		return err
	}
	c.ContainersConfig = containersConfig
	c.ImagesConfig = imagesConfig
	c.HabConfig = habConfig
	return nil
}

func (c *Config) GetContainerConfig(containerName string) (bases.HabContainerConfig, error) {

	for _, container := range c.ContainersConfig {
		if container.Name == containerName {
			return container, nil
		}
	}

	return bases.HabContainerConfig{}, errors.New("container not found")

}
