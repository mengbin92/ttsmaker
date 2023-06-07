package ttsmaker

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func (c *Client) GetVoiceList(ctx context.Context) (response ListResponse, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s?token=%s", listUrl, c.token), nil)
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

func (c *Client) GetVoiceListWithLang(ctx context.Context, lang string) (response ListResponse, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s?token=%s&language=%s", listUrl, c.token, lang), nil)
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
