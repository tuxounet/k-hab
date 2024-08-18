package config

type HabConfig = map[string]interface{}

type HabImageConfig struct {
	Name    string `yaml:"name"`
	Builder string `yaml:"builder"`
}

type HabContainerConfig struct {
	Name          string `yaml:"name"`
	Image         string `yaml:"image"`
	Shell         string `yaml:"shell"`
	Exec          string `yaml:"exec"`
	CloudInit     string `yaml:"cloud-init"`
	NetworkConfig string `yaml:"network-config"`
}
