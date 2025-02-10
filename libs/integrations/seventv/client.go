package seventv

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	baseURL = "https://7tv.io/v3"
)

type Client struct {
	client http.Client
}

type ClientOption func()

func NewClient() Client {
	return Client{
		client: http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (c *Client) TwitchUser(ctx context.Context, userID string) (Connection, error) {
	const endpoint = "/users/twitch/"

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+endpoint+userID, nil)
	if err != nil {
		return Connection{}, fmt.Errorf("new request with context: %w", err)
	}

	response, err := c.client.Do(request)
	if err != nil {
		return Connection{}, fmt.Errorf("do request: %w", err)
	}
	defer func() {
		_ = response.Body.Close()
	}()

	if response.StatusCode > 299 {
		return Connection{}, nil
	}

	var connection Connection

	if err = json.NewDecoder(response.Body).Decode(&connection); err != nil {
		return Connection{}, fmt.Errorf("decode response body: %w", err)
	}

	return connection, nil
}

func (c *Client) GlobalEmotes(ctx context.Context) (EmoteSet, error) {
	const endpoint = "/emote-sets/global"

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+endpoint, nil)
	if err != nil {
		return EmoteSet{}, fmt.Errorf("new request with context: %w", err)
	}

	response, err := c.client.Do(request)
	if err != nil {
		return EmoteSet{}, fmt.Errorf("do request: %w", err)
	}
	defer func() {
		_ = response.Body.Close()
	}()

	var globalEmotes EmoteSet

	if err = json.NewDecoder(response.Body).Decode(&globalEmotes); err != nil {
		return EmoteSet{}, fmt.Errorf("decode response body: %w", err)
	}

	return globalEmotes, nil
}
