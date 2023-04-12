package config

import (
	"github.com/caarlos0/env/v6"
)

type Configs struct {
	AppEnv               string `env:"APP_ENV" envDefault:"local"`
}

func Env() bool {
	return LoadEnv(output interface{})
}

func LoadEnv(output interface{}) bool {
	if e0 := env.Parse()
}