package bases

type HabConfig = map[string]interface{}

type HabBaseConfig struct {
	Name          string                 `yaml:"name"`
	Builder       string                 `yaml:"builder"`
	CloudInit     map[string]interface{} `yaml:"cloudinit"`
	NetworkConfig map[string]interface{} `yaml:"networkconfig"`
}
