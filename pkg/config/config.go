package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port     string   `yaml:"port"`
	Database Database `yaml:"database"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Migrate  bool   `yanl:"migrate"`
	SslMode  string `yaml:"sslmode"`
}

func ReadValue() *Config {
	var configs Config
	filename, _ := filepath.Abs("./config.yaml")
	yamlFile, _ := os.ReadFile(filename)
	yaml.Unmarshal(yamlFile, &configs)
	return &configs
}
