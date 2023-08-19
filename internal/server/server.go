package server

import (
	"context"
	"fmt"
	"grpc-server/configs"
	"grpc-server/internal/appctx"
	ecslogger "grpc-server/internal/logger"
	"net"
	"strconv"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type Server interface {
	Init()
}

type ServerOptions struct {
	Context context.Context
}

type server struct {
	grpcServer *grpc.Server
	ServerOptions
}

func NewServer(opt ServerOptions) Server {
	logger := appctx.FromContext(opt.Context)

	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		logging.WithFieldsFromContext(func(ctx context.Context) logging.Fields {
			return logging.Fields{"trace.id", ecslogger.TraceFromContext(ctx)}
		}),
		logging.WithDurationField(func(duration time.Duration) logging.Fields {
			return logging.Fields{"time.ms", float32(duration/1000) / 1000}
		}),
	}

	grpsServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		logging.UnaryServerInterceptor(ecslogger.InterceptorLogger(logger), opts...),
		func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
			fields := logging.ExtractFields(ctx)
			f := make([]zap.Field, 0, len(fields)/2)
			for i := fields.Iterator(); i.Next(); {
				k, v := i.At()
				f = append(f, zap.Any(k, v))
			}
			return handler(appctx.WithLogger(ctx, logger.With(f...)), req)
		},
		recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(func(a any) (err error) {
			return status.Errorf(codes.Unknown, "fail: %s", a)
		})),
	))

	return server{
		grpcServer:    grpsServer,
		ServerOptions: opt,
	}
}

func (s server) Init() {
	cfg := configs.GetConfig()

	port, err := strconv.Atoi(cfg.Server.Port)
	if err != nil {
		panic(err)
	}

	portGw, err := strconv.Atoi(cfg.Server.PortGw)
	if err != nil {
		panic(err)
	}

	reflection.Register(s.grpcServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	print(portGw)
	print(lis)
}
