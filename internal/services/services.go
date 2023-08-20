package services

import (
	"context"
	"grpc-server/configs"
	"grpc-server/internal/currency"
	"grpc-server/internal/services/provider"

	"grpc-server/internal/proto/gen/schema/services"
)

type ServiceServer struct {
	provider *provider.Provider
	services.UnimplementedCurrencyServiceServer
}

func New(provider *provider.Provider) *ServiceServer {
	return &ServiceServer{provider: provider}
}

func (s *ServiceServer) GetCurrency(ctx context.Context, req *services.CurrentRequest) (*services.CurrentResponse, error) {
	currencyFetch := currency.NewCurrencyFetch(s.provider.Client, configs.GetConfig().CurrencyFetch)
	serv := NewService(s.provider, currencyFetch)

	return serv.Run(ctx, req)
}
