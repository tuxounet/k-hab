package config

import (
	_ "embed"

	"github.com/tuxounet/k-hab/utils"
)

//go:embed templates/hab.yaml
var defaultHabConfig string

//go:embed templates/containers.yaml
var defaultContainersConfig string

//go:embed templates/images.yaml
var defaultImagesConfig string

type Config struct {
	scopeBase        string
	HabConfig        HabConfig
	ContainersConfig []HabContainerConfig
	ImagesConfig     []HabImageConfig
}

func NewConfig() *Config {
	return &Config{
		scopeBase: "Config",
	}
}

func (c *Config) Load(ctx *utils.ScopeContext) error {
	return ctx.Scope(c.scopeBase, "Load", func(ctx *utils.ScopeContext) {
		c.HabConfig = utils.LoadYamlFromString[HabConfig](ctx, defaultHabConfig)
		c.ContainersConfig = utils.LoadYamlFromString[[]HabContainerConfig](ctx, defaultContainersConfig)
		c.ImagesConfig = utils.LoadYamlFromString[[]HabImageConfig](ctx, defaultImagesConfig)
	})
}

func (c *Config) GetContainerConfig(ctx *utils.ScopeContext, containerName string) HabContainerConfig {
	return utils.ScopingWithReturn(ctx, c.scopeBase, "GetContainerConfig", func(ctx *utils.ScopeContext) HabContainerConfig {

		for _, container := range c.ContainersConfig {
			if container.Name == containerName {
				return container
			}
		}
		ctx.Must(ctx.Error("Container Not Found"))
		return HabContainerConfig{}
	})
}
