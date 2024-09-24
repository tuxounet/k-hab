package bases

import "encoding/base64"

type SetupFile struct {
	Config     map[string]string `yaml:"config"`
	Containers []SetupContainer  `yaml:"containers"`
}

type SetupContainer struct {
	Name     string                 `yaml:"name"`
	Base     string                 `yaml:"base"`
	Shell    string                 `yaml:"shell"`
	Network  map[string]interface{} `yaml:"network"`
	Proxy    map[string]interface{} `yaml:"proxy"`
	Deploy   string                 `yaml:"deploy"`
	Undeploy string                 `yaml:"undeploy"`
}

func (hcc *SetupContainer) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"name":     hcc.Name,
		"base":     hcc.Base,
		"shell":    hcc.Shell,
		"network":  hcc.Network,
		"proxy":    hcc.Proxy,
		"deploy":   base64.StdEncoding.EncodeToString([]byte(hcc.Deploy)),
		"undeploy": base64.StdEncoding.EncodeToString([]byte(hcc.Undeploy)),
	}
}
