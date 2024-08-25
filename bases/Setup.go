package bases

type SetupFile struct {
	Config     map[string]string `yaml:"config"`
	Containers []SetupContainer  `yaml:"containers"`
}

type SetupContainer struct {
	Name        string                 `yaml:"name"`
	Base        string                 `yaml:"base"`
	Shell       string                 `yaml:"shell"`
	Entry       string                 `yaml:"entry"`
	Network     map[string]interface{} `yaml:"network"`
	Proxy       map[string]interface{} `yaml:"proxy"`
	Provision   string                 `yaml:"provision"`
	Unprovision string                 `yaml:"unprovision"`
}

func (hcc *SetupContainer) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":        hcc.Name,
		"base":        hcc.Base,
		"shell":       hcc.Shell,
		"entry":       hcc.Entry,
		"network":     hcc.Network,
		"proxy":       hcc.Proxy,
		"provision":   hcc.Provision,
		"unprovision": hcc.Unprovision,
	}
}
