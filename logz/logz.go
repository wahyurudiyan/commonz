package logz

import (
	"context"

	"go.uber.org/zap"
)

type Logger interface {
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Debug(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
	Fatal(ctx context.Context, msg string, fields ...zap.Field)
	Panic(ctx context.Context, msg string, fields ...zap.Field)
	Warning(ctx context.Context, msg string, fields ...zap.Field)
	DevelopmentPanic(ctx context.Context, msg string, fields ...zap.Field)
}

func NewDefault() *defaultLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(InitDefaultErr)
	}

	return &defaultLogger{
		logger: logger,
	}
}
