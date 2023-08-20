package server

import (
	"context"
	"fmt"
	"grpc-server/configs"
	"grpc-server/internal/appctx"
	ecslogger "grpc-server/internal/logger"
	"grpc-server/internal/services/provider"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"grpc-server/internal/proto/gen/schema/services"
	serv "grpc-server/internal/services"

	"github.com/fullstorydev/grpcui/standalone"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type Server interface {
	Init()
}

type ServerOptions struct {
	Context context.Context
	Client  http.Client
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

	provider := provider.NewProvider(s.Client)

	services.RegisterCurrencyServiceServer(s.grpcServer, serv.New(provider))

	healtcheck := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s.grpcServer, healtcheck)

	go func() {
		fmt.Println("server started - port:", port)
		if err := s.grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	go func() {
		conn, err := grpc.DialContext(
			s.Context,
			fmt.Sprintf(":%d", port),
			grpc.WithBlock(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			panic(err)
		}

		gmux := runtime.NewServeMux()
		err = services.RegisterCurrencyServiceHandler(s.Context, gmux, conn)
		if err != nil {
			panic(err)
		}

		gwServer := &http.Server{
			Addr:    fmt.Sprintf(":%d", portGw),
			Handler: gmux,
		}

		fmt.Println("gRPC Gateway server on: ", portGw)
		log.Fatalln(gwServer.ListenAndServe())
	}()

	time.Sleep(time.Second)
	configureGrpcUI(context.Background(), http.DefaultServeMux, port)
}

func configureGrpcUI(ctx context.Context, mux *http.ServeMux, port int) {
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	handler, err := standalone.HandlerViaReflection(ctx, conn, fmt.Sprintf("%s:%d", filepath.Base(os.Args[0])))
	if err != nil {
		panic(err)
	}

	mux.Handle("/ui/", http.StripPrefix("/ui", handler))
	fmt.Println("grpcUI started - http://127.0.0.1:9071/ui/")

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", 9071))
	if err != nil {
		panic(err)
	}

	err = http.Serve(listen, mux)
	if err != nil {
		panic(err)
	}
}
