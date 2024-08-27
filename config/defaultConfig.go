package config

import (
	_ "embed"
)

//go:embed default.config.yaml
var DefaultConfig string

//go:embed default.setup.yaml
var DefaultSetup string
