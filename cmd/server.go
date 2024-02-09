package main

import (
	"flag"
	"fmt"

	config "github.com/BenjixD/bengi-lol-bot/utils"
	"go.uber.org/zap"
)

var (
	cfg    *config.BengiLolBotConfig
	logger *zap.Logger

	host     string
	httpPort int
	grpcPort int
)

func init() {
	var err error
	cfg, err = config.Init()
	if err != nil {
		panic(fmt.Sprintf("Could not read configuration: %v", err))
	}

	switch env := cfg.Environment; env {
	case config.Development:
		logger, err = zap.NewDevelopment()
	case config.Staging:
		logger, err = zap.NewDevelopment()
	case config.Production:
		logger, err = zap.NewProduction()
	default:
		panic("cannot initialize logger, unknown environment found")
	}
	if err != nil {
		panic("failed to initialize logger")
	}

	flag.StringVar(&host, "host", "localhost", "server hostname")
	flag.IntVar(&httpPort, "http-port", 8080, "server http port")
	flag.IntVar(&grpcPort, "grpc-port", 8081, "server grpc port")
	flag.Parse()
}

func main() {
	logger.Debug(fmt.Sprintf("%v", cfg.Environment))
	logger.Debug(fmt.Sprintf("%v", cfg.ApiKey))
}
