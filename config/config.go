package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

var (
	config Config
)

type Config struct {
	PortNo int64 `yaml:"port_no"`
}

func init() {

	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(configFile, &config); err != nil {
		panic(err)
	}

}

func Get() Config {
	return config
}
