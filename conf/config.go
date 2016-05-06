package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type config struct {
	file string
}

// NewConfig constructor
func NewConfig(file string) *config {
	return &config{file}
}

// Load config from file
func (this *config) Load() (*Config, error) {
	data, err := ioutil.ReadFile(this.file)
	if err != nil {
		return nil, err
	}

	var config *Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

type Config struct {
	Debug  bool   `yaml:"debug"`
	Server Server `yaml:"server"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
