package logger

import (
	"context"
	"grpc-server/internal/constants"
	"os"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/metadata"
)

func NewEcsLogger(ctx context.Context) *zap.Logger {
	config := ecszap.EncoderConfig{
		EncodeDuration: zapcore.MillisDurationEncoder,
	}

	core := ecszap.NewCore(
		config,
		os.Stdout,
		zap.DebugLevel,
	)

	logger := zap.New(core, zap.AddCallerSkip(1))

	defer logger.Sync()

	return logger
}

func InterceptorLogger(logger *zap.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		f := make([]zap.Field, 0, len(fields)/2)
		i := logging.Fields(fields).Iterator()
		for i.Next() {
			k, v := i.At()
			f = append(f, zap.Any(k, v))
		}
		l := logger.WithOptions(zap.AddCallerSkip(1)).With(f...)

		switch lvl {
		case logging.LevelDebug:
			l.Debug(msg)
		case logging.LevelInfo:
			l.Info(msg)
		case logging.LevelWarn:
			l.Warn(msg)
		case logging.LevelError:
			l.Error(msg)
		default:
			panic("unknown log level")
		}
	})
}

func TraceFromContext(ctx context.Context) string {
	traceId := metadata.ValueFromIncomingContext(ctx, constants.GetHeaders.TraceId)
	if len(traceId) == 0 {
		return uuid.NewString()
	}

	return traceId[0]
}
