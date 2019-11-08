package cacoo

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Client struct {
	client *http.Client
	apiKey string
	debug  bool
}

func NewClient(apiKey string, debug bool) *Client {
	return &Client{
		client: http.DefaultClient,
		apiKey: apiKey,
		debug:  debug,
	}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	req2 := req.Clone(req.Context())
	if c.apiKey != "" {
		q := req2.URL.Query()
		q.Add("apiKey", c.apiKey)
		req2.URL.RawQuery = q.Encode()
	}
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
		if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
			return fmt.Errorf("request to Cacoo API was not successful: %s", resp.Status)
		}
		var r io.Reader = resp.Body
		if c.debug {
			r = io.TeeReader(r, os.Stdout)
		}
		err = xml.NewDecoder(r).Decode(v)
		if err != nil {
			return err
		}
		return nil
	}
}
