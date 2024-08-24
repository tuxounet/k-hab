package bases

type HabConfig = map[string]interface{}

type HabImageConfig struct {
	Name    string `yaml:"name"`
	Builder string `yaml:"builder"`
}

type HabContainerConfig struct {
	Name          string                 `yaml:"name"`
	Image         string                 `yaml:"image"`
	Shell         string                 `yaml:"shell"`
	Entry         string                 `yaml:"entry"`
	Network       map[string]interface{} `yaml:"network"`
	Proxy         map[string]interface{} `yaml:"proxy"`
	CloudInit     string                 `yaml:"cloud-init"`
	NetworkConfig string                 `yaml:"network-config"`
}

func (hcc *HabContainerConfig) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":           hcc.Name,
		"image":          hcc.Image,
		"shell":          hcc.Shell,
		"entry":          hcc.Entry,
		"network":        hcc.Network,
		"proxy":          hcc.Proxy,
		"cloud-init":     hcc.CloudInit,
		"network-config": hcc.NetworkConfig,
	}
}
