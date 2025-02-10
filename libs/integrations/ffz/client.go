package ffz

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	baseURL = "https://api.frankerfacez.com/v1"
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

func (c *Client) RoomEmoteSetCollection(ctx context.Context, roomID string) (EmoteSetCollection, error) {
	const endpoint = "/room/id/"

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+endpoint+roomID, nil)
	if err != nil {
		return EmoteSetCollection{}, fmt.Errorf("new request with context: %w", err)
	}

	response, err := c.client.Do(request)
	if err != nil {
		return EmoteSetCollection{}, fmt.Errorf("do request: %w", err)
	}
	defer func() {
		_ = response.Body.Close()
	}()

	var collection EmoteSetCollection

	if err = json.NewDecoder(response.Body).Decode(&collection); err != nil {
		return EmoteSetCollection{}, fmt.Errorf("decode response body: %w", err)
	}

	return collection, nil
}

func (c *Client) GlobalEmoteSetCollection(ctx context.Context) (EmoteSetCollection, error) {
	const endpoint = "/set/global"

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+endpoint, nil)
	if err != nil {
		return EmoteSetCollection{}, fmt.Errorf("new request with context: %w", err)
	}

	response, err := c.client.Do(request)
	if err != nil {
		return EmoteSetCollection{}, fmt.Errorf("do request: %w", err)
	}
	defer func() {
		_ = response.Body.Close()
	}()

	var collection EmoteSetCollection

	if err = json.NewDecoder(response.Body).Decode(&collection); err != nil {
		return EmoteSetCollection{}, fmt.Errorf("decode response body: %w", err)
	}

	return collection, nil
}
