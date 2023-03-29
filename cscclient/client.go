package cscclient

import (
	"net/http"
)

const BaseURL = "https://apis.cscglobal.com/dbs/api/v2"

type Client struct {
	APIKey       string
	BearerToken  string
	HTTPClient   *http.Client
}

func NewClient(apiKey, bearerToken string) *Client {
	return &Client{
		APIKey:      apiKey,
		BearerToken: bearerToken,
		HTTPClient:  &http.Client{},
	}
}
