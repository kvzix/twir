package ffz

import (
	"context"
	"fmt"

	"github.com/satont/twir/apps/emotes-cacher/internal/emote"
	"github.com/twirapp/twir/libs/integrations/ffz"
)

type Provider struct {
	client ffz.Client
}

var _ emote.Provider = (*Provider)(nil)

func NewProvider(client ffz.Client) *Provider {
	return &Provider{
		client: client,
	}
}

func (p *Provider) Global(ctx context.Context) ([]emote.Emote, error) {
	collection, err := p.client.GlobalEmoteSetCollection(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetch global emote set collection: %w", err)
	}

	return p.collectionToEmotes(collection), nil
}

func (p *Provider) Channel(ctx context.Context, channelID string) ([]emote.Emote, error) {
	collection, err := p.client.RoomEmoteSetCollection(ctx, channelID)
	if err != nil {
		return nil, fmt.Errorf("fetch room emote set collection: %w", err)
	}

	return p.collectionToEmotes(collection), nil
}

func (p *Provider) collectionToEmotes(collection ffz.EmoteSetCollection) []emote.Emote {
	emotesIDs := make([]emote.Emote, 0)

	for _, emoteSet := range collection.EmoteSets {
		for _, globalEmote := range emoteSet.Emotes {
			emotesIDs = append(emotesIDs, emote.Emote{
				ID:       globalEmote.ID,
				Name:     globalEmote.Name,
				Provider: "ffz",
			})
		}
	}

	return emotesIDs
}
