package config

import (
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
			RedirectURL  string `yaml:"redirect_url"`
		} `yaml:"oauth2_config"`
	} `yaml:"google"`

	JWTExpiredHours int64 `yaml:"jwt_expired_hours"`

	Mail struct {
		SenderMail     string `yaml:"sender_mail"`
		SenderPassword string `yaml:"sender_password"`
	} `yaml:"mail"`
}

func init() {
	configFile, err := os.ReadFile(configFilePath)
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
