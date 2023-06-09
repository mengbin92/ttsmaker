package ttsmaker

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
)

type Client struct {
	token  string
	client *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		token:  token,
		client: &http.Client{},
	}
}

// sendRequest sends an HTTP request to a given URL and handles the response
func (c *Client) sendRequest(req *http.Request, v any) error {
	res, err := c.client.Do(req)
	if err != nil {
		return errors.Wrap(err, "send ChatCompletion request error")
	}
	defer res.Body.Close()

	// Handle the response from ttsmaker
	if v != nil {
		if err := sonic.ConfigDefault.NewDecoder(res.Body).Decode(v); err != nil {
			return errors.Wrap(err, "unmarshal ttsmaker Response error")
		}
	}
	return nil
}
