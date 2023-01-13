package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"database"`
}

var (
	cfg *Config = nil
)

func Initialize(configPath string) {
	loadConfig(configPath + "/default.yaml")
	loadConfig(configPath + "/override.yaml")
}

func loadConfig(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	data := yaml.NewDecoder(file)
	if err := data.Decode(&cfg); err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	return cfg
}
