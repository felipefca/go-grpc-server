package appctx

import (
	"context"

	"go.uber.org/zap"
)

var (
	LoggerKey = "logger"
)

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, LoggerKey, logger)
}

func FromContext(ctx context.Context) *zap.Logger {
	logger := ctx.Value(LoggerKey).(*zap.Logger)
	return logger
}
