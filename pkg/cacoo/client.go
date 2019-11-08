package cacoo

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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

func (c *Client) do(req *http.Request, err error) func(interface{}) error {
	return func(v interface{}) error {
		if err != nil {
			return err
		}
		resp, err := c.Do(req)
		if err != nil {
			return err
		}
		defer func() {
			_, err = io.Copy(ioutil.Discard, resp.Body)
			if err != nil {
			}
			err = resp.Body.Close()
			if err != nil {

			}
		}()
		err = xml.NewDecoder(io.TeeReader(resp.Body, os.Stderr)).Decode(v)
		if err != nil {
			return err
		}
		return nil
	}
}
