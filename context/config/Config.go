package config

import (
	_ "embed"

	"github.com/tuxounet/k-hab/bases"
)

type Config struct {
	values map[string]string
	log    bases.ILogger
}

func NewConfig(logger bases.ILogger, defaultConfig map[string]string) *Config {
	return &Config{
		values: defaultConfig,
		log:    logger.CreateSubLogger("Setup", logger),
	}
}

func (c *Config) SetConfigValue(key string, value string) error {

	c.values[key] = value

	c.log.TraceF("Config value %s set to %s", key, value)
	return nil
}

func (c *Config) GetValue(key string) string {
	c.log.TraceF("Config value %s requested", key)
	return c.values[key]
}

func (c *Config) GetCurrent() map[string]string {

	return c.values
}
