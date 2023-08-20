package main

import (
	"context"
	"grpc-server/internal/appctx"
	"grpc-server/internal/logger"
	"grpc-server/internal/server"
	"net/http"
)

func main() {
	ctx := context.Background()

	client := http.Client{}

	logger := logger.NewEcsLogger(ctx)
	ctx = appctx.WithLogger(ctx, logger)

	s := server.NewServer(server.ServerOptions{
		Context: ctx,
		Client:  client,
	})
	s.Init()
}
