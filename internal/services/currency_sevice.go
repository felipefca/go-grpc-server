package services

import (
	"context"
	"grpc-server/internal/appctx"
	"grpc-server/internal/currency"
	"grpc-server/internal/proto/gen/schema/common"
	schema_currency "grpc-server/internal/proto/gen/schema/currency"
	"grpc-server/internal/proto/gen/schema/services"
	"grpc-server/internal/services/provider"
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type CurrencyServer struct {
	provider *provider.Provider
	currency currency.CurrencyInterface
}

func NewService(provider *provider.Provider, currency currency.CurrencyInterface) *CurrencyServer {
	return &CurrencyServer{
		provider: provider,
		currency: currency,
	}
}

func (c CurrencyServer) Run(ctx context.Context, req *services.CurrentRequest) (*services.CurrentResponse, error) {
	logger := appctx.FromContext(ctx)
	logger.Info("Starting Run")

	response, err := c.currency.GetCotation(ctx, req.Code, req.CodeIn)
	if err != nil {
		return nil, err
	}

	date := timestamppb.New(parseDate(response.Date))

	output := &services.CurrentResponse{
		Currency: &schema_currency.Currency{
			Code:   response.Code,
			CodeIn: response.CodeIn,
			Name:   response.Name,
			Date:   date,
			Price: &common.Price{
				High:  parseFloat64(response.High),
				Low:   parseFloat64(response.Low),
				Value: parseFloat64(response.Value),
			},
		},
	}

	logger.Info("Successfully Run")
	return output, nil
}

func parseFloat64(value string) float64 {
	f, _ := strconv.ParseFloat(value, 64)
	return f
}

func parseDate(date string) time.Time {
	parsedTime, _ := time.Parse("2006-01-02 15:04:05", date)
	return parsedTime
}
