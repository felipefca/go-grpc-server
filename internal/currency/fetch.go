package currency

import (
	"context"
	"encoding/json"
	"fmt"
	"grpc-server/configs"
	"grpc-server/internal/appctx"
	"io/ioutil"
	"net/http"
)

type CurrencyInterface interface {
	GetCotation(ctx context.Context, code string, codeIn string) (*Response, error)
}

type Currency struct {
	Client http.Client
	Config configs.CurrencyFetch
}

func NewCurrencyFetch(client http.Client, config configs.CurrencyFetch) Currency {
	return Currency{
		Client: client,
		Config: config,
	}
}

func (c Currency) GetCotation(ctx context.Context, code string, codeIn string) (*Response, error) {
	var response Response
	logger := appctx.FromContext(ctx)

	url := fmt.Sprintf("%s%s-%s", c.Config.UrlLastContation, code, codeIn)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Fail because : %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	logger.Info("Success to GetCotation")
	return &response, nil
}
