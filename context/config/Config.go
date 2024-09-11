package config

import (
	_ "embed"

	"github.com/tuxounet/k-hab/bases"
)

type Config struct {
	values map[string]string
	log    bases.ILogger
}

func NewConfig(defaultConfig map[string]string) *Config {
	return &Config{
		values: defaultConfig,
	}
}
func (c *Config) SetLogger(logger bases.ILogger) {
	c.log = logger.CreateSubLogger("Setup", logger)
}

func (c *Config) SetConfigValue(key string, value string) error {

	c.values[key] = value
	if c.log != nil {
		c.log.TraceF("Config value %s set to %s", key, value)
	}
	return nil
}

func (c *Config) GetValue(key string) string {
	if c.log != nil {
		c.log.TraceF("Config value %s requested", key)
	}

	return c.values[key]
}

func (c *Config) GetCurrent() map[string]string {

	return c.values
}
