package env

import (
	env_v9 "github.com/caarlos0/env/v9"
)

type Env struct {
	DB struct {
		User     string `env:"DB_USER" envDefault:"postgres"`
		Password string `env:"DB_PASSWORD" envDefault:"postgres"`
		Host     string `env:"DB_HOST" envDefault:"127.0.0.1"`
		Port     int64  `env:"DB_PORT" envDefault:"5432"`
		Database string `env:"DB_DATABASE" envDefault:"kikumemodb"`
	}
	Http struct {
		Port int64 `env:"HTTP_PORT" envDefault:"8081"`
	}
	AllowOrigins []string `env:"ALLOW_ORIGINS" envDefault:"[]"`
}

var (
	env Env
)

func init() {

	if err := env_v9.Parse(&env); err != nil {
		panic(err)
	}

}

func Get() Env {
	return env
}
