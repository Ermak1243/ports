package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	InPortsCount  int    `yaml:"in_ports_count"`
	OutPortsCount int    `yaml:"out_ports_count"`
	ServerPort    string `yaml:"server_port"`
}

func NewConfig(path string) *Config {
	return MustLoadPath(path) // Load configuration using MustLoadPath function
}

func MustLoadPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("config path is empty: " + err.Error())
	}

	return &cfg
}
