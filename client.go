package ttsmaker

import (
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
)

type Client struct {
	token      string
	HttpClient *http.Client
}

func NewClinet(token string) *Client {
	return &Client{
		token:      token,
		HttpClient: &http.Client{},
	}
}

func (c *Client) sendRequest(req *http.Request, v any) error {
	res, err := c.HttpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "send ChatCompletion request error")
	}
	defer res.Body.Close()

	if v != nil {
		if err := sonic.ConfigDefault.NewDecoder(res.Body).Decode(v); err != nil {
			return errors.Wrap(err, "unmarshal ttsmaker Response error")
		}
	}
	return nil
}
