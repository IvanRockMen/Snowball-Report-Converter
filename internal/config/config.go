package config

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

const defaultPath = "config.yaml"

type Config struct {
	Plugins map[string]string `yaml:"plugins"`
}

func LoadConfig() (*Config, error) {
	cfg := new(Config)
	path := ""
	p := flag.String("config", "", "config path for work")
	if p == nil {
		path = defaultPath
	} else {
		path = *p
	}
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(f, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
