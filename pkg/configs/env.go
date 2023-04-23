package configs

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type ConfigEnv struct {
	AppEnv     string `env:"APP_ENV" envDefault:"local"`
	AppPort    int16  `env:"APP_PORT" envDefault:"80"`
	RedisHost  string `env:"REDIS_HOST" envDefault:"localhost:6379"`
	RedisPower string `env:"REDIS_POWER" envDefault:"ON"`
}

var configEnv = &ConfigEnv{}

func Env() bool {
	// Load file env
	e0 := godotenv.Load("env/.env")
	if e0 != nil {
		fmt.Printf("Load env config failure. Err: %s", e0)
	}

	fmt.Printf("Load env config successfully")

	// Load env to Configs
	if e1 := env.Parse(configEnv); e1 != nil {
		fmt.Printf("Load env to configEnv error. Err: %s", e1)
	}

	return true
}

func AppEnv() string {
	return configEnv.AppEnv
}

func AppPort() int16 {
	return configEnv.AppPort
}

func RedisHost() string {
	return configEnv.RedisHost
}

func RedisPower() string {
	return configEnv.RedisPower
}
