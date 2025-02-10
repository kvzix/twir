package emotes

import (
	"context"
	"time"
)

type CacheRepository interface {
	SetGlobalMany(ctx context.Context, emotes []SetEmoteInput, expiration time.Duration) error
	SetChannelMany(ctx context.Context, channelID string, emotes []SetEmoteInput, expiration time.Duration) error
}

type SetEmoteInput struct {
	ID       string
	Name     string
	Provider string
}
