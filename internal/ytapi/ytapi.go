package ytapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	apiKeys    []key
}

type key struct {
	apiKey       string
	quotaReached bool
}

func NewClient(apiKeys []string) Client {
	var c Client
	c.httpClient = http.Client{
		Timeout: time.Minute,
	}
	c.apiKeys = make([]key, 0)
	for _, k := range apiKeys {
		c.apiKeys = append(c.apiKeys, key{k, false})
	}
	return c
}
