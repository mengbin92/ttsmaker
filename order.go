package ttsmaker

import (
	"bytes"
	"context"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
)

// Create a TTS order. The Create TTS Order API allows you to submit a TTS order and receive a temporary URL to access the generated speech file.
// Please note that the API does not return the speech file itself, but instead provides a URL to a temporary file that will be deleted after 2 hours. 
// You can use this URL to play or download the speech file as part of your business process.
func (c *Client) CreateOrder(ctx context.Context, request *OrderRequest) (resonse OrderResponse, err error) {
	payload, err := sonic.Marshal(request)
	if err != nil {
		errors.Wrap(err, "marshal OrderRequest error")
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, orderUrl, bytes.NewBuffer(payload))
	if err != nil {
		errors.Wrap(err, "New http request for create order error")
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "TTSMaker API Client")

	err = c.sendRequest(req, &resonse)
	return
}
