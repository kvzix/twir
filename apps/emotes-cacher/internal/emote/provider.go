package emote

import (
	"context"
)

type (
	Emote struct {
		ID   string
		Name string
		// Provider is a name of the service that provides this emote.
		Provider string
	}

	Provider interface {
		GlobalEmotes(ctx context.Context) ([]Emote, error)
		ChannelEmotes(ctx context.Context, channelID string) ([]Emote, error)
	}
)
