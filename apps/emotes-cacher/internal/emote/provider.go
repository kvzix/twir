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
		Global(ctx context.Context) ([]Emote, error)
		Channel(ctx context.Context, channelID string) ([]Emote, error)
	}
)
