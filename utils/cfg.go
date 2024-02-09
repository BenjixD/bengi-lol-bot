package config

import (
	"github.com/spf13/viper"
)

type Env string

const (
	Production  Env = "production"
	Staging     Env = "staging"
	Development Env = "development"
)

type BengiLolBotConfig struct {
	Environment Env    `mapstructure:"ENVIRONMENT"`
	ApiKey      string `mapstructure:"API_KEY"`
}

func Init() (config *BengiLolBotConfig, err error) {
	v := viper.New()
	// Note: Unmarshal needs to bind envs to read from envar
	// https://github.com/spf13/viper/issues/761
	v.SetDefault("ENVIRONMENT", Development)
	v.SetDefault("API_KEY", "")
	v.AutomaticEnv()

	err = v.Unmarshal(&config)
	return config, err
}
