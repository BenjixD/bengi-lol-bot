package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/BenjixD/bengi-lol-bot/svc"
	config "github.com/BenjixD/bengi-lol-bot/utils/config"
	"github.com/BenjixD/bengi-lol-bot/utils/log"
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

	logger, err = log.NewLogger(cfg.Environment)
	if err != nil {
		panic("failed to initialize logger")
	}

	flag.StringVar(&host, "host", "localhost", "server hostname")
	flag.IntVar(&httpPort, "http-port", 8080, "server http port")
	flag.IntVar(&grpcPort, "grpc-port", 8081, "server grpc port")
	flag.Parse()
}

func main() {
	riotClient, _ := svc.NewRiotApiClient(cfg.RiotApiKey)
	res, err := riotClient.GetUserFromRiotID(context.TODO(), &svc.GetUserFromRiotIDRequest{
		GameName: "alierujah",
		TagLine:  "na1",
	})
	fmt.Println(res)
	fmt.Println(err)
}
