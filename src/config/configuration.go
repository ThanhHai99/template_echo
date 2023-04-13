package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Configs struct {
	AppEnv  string `env:"APP_ENV" envDefault:"local"`
	AppPort int16  `env:"APP_PORT" envDefault:"80"`
}

var configs = &Configs{}

func LoadEnv() bool {
	// Load file env
	e0 := godotenv.Load("env/.env")
	if e0 != nil {
		fmt.Printf("Load env config failure. Err: %s", e0)
	}

	fmt.Printf("Load env config successfully")

	// Load env to Configs
	if e1 := env.Parse(configs); e1 != nil {
		fmt.Printf("Load env to configs error. Err: %s", e1)
	}

	return true
}

func AppEnv() string {
	return configs.AppEnv
}

func AppPort() int16 {
	return configs.AppPort
}
