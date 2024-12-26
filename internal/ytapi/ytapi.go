package ytapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	apiKeys    []string
}

func NewClient(apiKeys []string) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		apiKeys: apiKeys,
	}
}
