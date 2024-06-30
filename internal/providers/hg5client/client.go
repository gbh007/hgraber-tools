package hg5client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client

	addr  string
	token string
}

func New(addr, token string) *Client {
	return &Client{
		client: &http.Client{
			Transport: http.DefaultTransport,
			Timeout:   time.Minute,
		},

		addr:  addr,
		token: token,
	}
}

func (c *Client) ImportArchive(ctx context.Context, body io.Reader) error {
	request, err := http.NewRequestWithContext(
		ctx, http.MethodPost,
		c.addr+"/api/system/import/archive", // Примечание, это очень плохой способ, но смысла делать лучше тут нет
		body,
	)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}

	request.Header.Set("X-HG-Token", c.token)
	request.Header.Set("Content-Type", "application/octet-stream")

	response, err := c.client.Do(request)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("unsuccess response (%d): %s", response.StatusCode, string(data))
	}

	return nil
}
