package provider

import "net/http"

type Provider struct {
	Client http.Client
}

func NewProvider(client http.Client) *Provider {
	c := &Provider{Client: client}
	return c
}
