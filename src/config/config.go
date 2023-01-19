package config

import (
	"os"
	"strconv"

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

func SetConfig(config *Config) {
	cfg = config
}

func GetDatabaseAddr() string {
	return cfg.Database.Host + ":" + strconv.Itoa(cfg.Database.Port)
}

func GetServerAddr() string {
	return cfg.Server.Host + ":" + strconv.Itoa(cfg.Server.Port)
}

func GetServerUrl(path string) string {
	return "http://" + GetServerAddr() + "/" + path
}
