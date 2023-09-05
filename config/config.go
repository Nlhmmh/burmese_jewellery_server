package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	config Config
)

type Config struct {
	Google struct {
		Oauth2Config struct {
			ClientID     string `yaml:"client_id"`
			ClientSecret string `yaml:"client_secret"`
			RedirectURL  string `env:"redirect_url"`
		} `yaml:"oauth2_config"`
	} `yaml:"google"`
}

func init() {

	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(configFile, &config); err != nil {
		panic(err)
	}

	fmt.Println(config)

}

func Get() Config {
	return config
}
