package config

type ConfigData struct {
	Database struct {
		Type             string `yaml:"type"`
		Target           string `yaml:"target"`
		LogRetentionDays int    `yaml:"logRetentionDays"`
	} `yaml:"database"`
	Port int `yaml:"port"`
}
