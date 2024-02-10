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
	Environment   Env    `mapstructure:"ENVIRONMENT"`
	DiscordApiKey string `mapstructure:"DISCORD_API_KEY"`
	AiApiKey      string `mapstructure:"AI_API_KEY"`
	RiotApiKey    string `mapstructure:"RIOT_API_KEY"`
}

func Init() (config *BengiLolBotConfig, err error) {
	v := viper.New()
	// Note: Unmarshal needs to bind envs to read from envar
	// https://github.com/spf13/viper/issues/761
	v.SetDefault("ENVIRONMENT", Development)
	v.SetDefault("DISCORD_API_KEY", "")
	v.SetDefault("AI_API_KEY", "")
	v.SetDefault("RIOT_API_KEY", "")
	v.AutomaticEnv()

	err = v.Unmarshal(&config)
	return config, err
}
