package log

import (
	"context"
	"fmt"

	"github.com/BenjixD/bengi-lol-bot/utils/config"
	"go.uber.org/zap"
)

type Key struct{}

func NewLogger(env config.Env) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error = nil
	switch env {
	case config.Development:
		logger, err = zap.NewDevelopment()
	case config.Staging:
		logger, err = zap.NewDevelopment()
	case config.Production:
		logger, err = zap.NewProduction()
	}
	return logger, err
}

func WithContext(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, Key{}, logger)
}

func FromContext(ctx context.Context) (*zap.Logger, error) {
	logger := ctx.Value(Key{})
	if l, ok := logger.(*zap.Logger); ok {
		return l, nil
	}
	return nil, fmt.Errorf("logger not initialized")
}
