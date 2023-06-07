package ttsmaker

import (
	"bytes"
	"context"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
)

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
