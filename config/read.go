package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

func ReadConfig(path string) (Config, error) {
	var cfg Config
	cfgFile, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}

	return cfg, yaml.Unmarshal(cfgFile, &cfg)
}

func MustReadConfig(path string) Config {
	cfg, err := ReadConfig(path)
	if err != nil {
		panic(err)
	}
	return cfg
}
