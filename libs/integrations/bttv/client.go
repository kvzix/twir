package bttv

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	baseURL = "https://api.betterttv.net/3"
)

type Client struct {
	client http.Client
}

func NewClient() Client {
	return Client{
		client: http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (c *Client) TwitchUser(ctx context.Context, userID string) (User, error) {
	const endpoint = "/cached/users/twitch/"

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+endpoint+userID, nil)
	if err != nil {
		return User{}, fmt.Errorf("new request with context: %w", err)
	}

	response, err := c.client.Do(request)
	if err != nil {
		return User{}, fmt.Errorf("do request: %w", err)
	}
	defer func() {
		_ = response.Body.Close()
	}()

	var user User

	if err = json.NewDecoder(response.Body).Decode(&user); err != nil {
		return User{}, fmt.Errorf("decode response body: %w", err)
	}

	return user, nil
}

func (c *Client) GlobalEmotes(ctx context.Context) ([]Emote, error) {
	const endpoint = "/cached/emotes/global"

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("new request with context: %w", err)
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer func() {
		_ = response.Body.Close()
	}()

	var globalEmotes []Emote

	if err = json.NewDecoder(response.Body).Decode(&globalEmotes); err != nil {
		return nil, fmt.Errorf("decode response body: %w", err)
	}

	return globalEmotes, nil
}
