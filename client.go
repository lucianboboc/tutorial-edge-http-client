package client

import "net/http"

var (
	DefaultAPIURL = "https://pokeapi.co"
)

type Client struct {
	apiURL     string
	httpClient *http.Client
}

type Option = func(c *Client)

func NewClient(options ...Option) *Client {
	client := Client{
		apiURL:     DefaultAPIURL,
		httpClient: http.DefaultClient,
	}

	for _, o := range options {
		o(&client)
	}

	return &client
}

func WithAPIURL(url string) Option {
	return func(c *Client) {
		c.apiURL = url
	}
}

func WithHttpClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
