package hg4client

import (
	"bytes"
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

func (c *Client) BookArchive(ctx context.Context, id int) (io.Reader, error) {
	request, err := http.NewRequestWithContext(
		ctx, http.MethodGet,
		fmt.Sprintf("%s/api/book/download?id=%d", c.addr, id), // Примечание, это очень плохой способ, но смысла делать лучше тут нет
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	request.Header.Set("X-Token", c.token)

	response, err := c.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("unsuccess response (%d): %s", response.StatusCode, string(data))
	}

	return bytes.NewReader(data), nil
}
