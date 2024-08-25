package bases

type HabConfig = map[string]interface{}

type HabBaseConfig struct {
	Name          string                 `yaml:"name"`
	Builder       string                 `yaml:"builder"`
	CloudInit     map[string]interface{} `yaml:"cloudinit"`
	NetworkConfig map[string]interface{} `yaml:"networkconfig"`
}

type HabContainerConfig struct {
	Name    string                 `yaml:"name"`
	Base    string                 `yaml:"base"`
	Shell   string                 `yaml:"shell"`
	Entry   string                 `yaml:"entry"`
	Network map[string]interface{} `yaml:"network"`
	Proxy   map[string]interface{} `yaml:"proxy"`
}

func (hcc *HabContainerConfig) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":    hcc.Name,
		"base":    hcc.Base,
		"shell":   hcc.Shell,
		"entry":   hcc.Entry,
		"network": hcc.Network,
		"proxy":   hcc.Proxy,
	}
}
