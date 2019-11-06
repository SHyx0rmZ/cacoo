package cacoo

import (
	"net/http"
)

type Client struct {
	client *http.Client
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		client: http.DefaultClient,
		apiKey: apiKey,
	}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req2 := req.Clone(req.Context())
	q := req2.URL.Query()
	q.Add("apiKey", c.apiKey)
	req2.URL.RawQuery = q.Encode()
	return c.client.Do(req2)
}
