package config

import (
	_ "embed"
	"errors"

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

func (c *Config) GetContainerConfig(ctx *utils.ScopeContext, containerName string) map[string]interface{} {
	return utils.ScopingWithReturnOnly(ctx, c.scopeBase, "GetContainerConfig", func(ctx *utils.ScopeContext) map[string]interface{} {

		for _, container := range c.ContainersConfig {
			if container.(map[string]interface{})["name"] == containerName {
				return container.(map[string]interface{})
			}
		}
		ctx.Must(errors.New("Container Not Found"))
		return nil
	})
}
