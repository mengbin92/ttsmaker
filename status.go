package ttsmaker

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func (c *Client) GetTokenStatus(ctx context.Context) (response StatusResponse, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s?token=%s", statusUrl, c.token), nil)
	if err != nil {
		errors.Wrap(err, "create get voice list request error")
		return
	}

	if err = c.sendRequest(req, &response); err != nil {
		errors.Wrap(err, "get voice list from ttsmaker error")
		return
	}
	return
}
