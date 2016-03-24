package conf

import (
	"github.com/jinzhu/configor"
)

var (
	config *Config
)

func init() {
	configor.Load(&config, "config.yaml")
}

func GetConfig() *Config {
	return config
}

type Config struct {
	Debug  bool   `yaml:"debug"`
	Server Server `yaml:"server"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
