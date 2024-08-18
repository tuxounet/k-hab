package config

type HabConfig = map[string]interface{}
type HabContainerConfig = interface{}

type HabImageConfig struct {
	Name    string `yaml:"name"`
	Builder string `yaml:"builder"`
}
