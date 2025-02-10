package seventv

import (
	"context"
	"fmt"

	"github.com/satont/twir/apps/emotes-cacher/internal/emote"
	"github.com/twirapp/twir/libs/integrations/seventv"
)

type Provider struct {
	client seventv.Client
}

var _ emote.Provider = (*Provider)(nil)

func NewProvider(client seventv.Client) *Provider {
	return &Provider{
		client: client,
	}
}

func (p *Provider) Global(ctx context.Context) ([]emote.Emote, error) {
	emotes, err := p.client.GlobalEmotes(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetch global emotes: %w", err)
	}

	return p.emoteSetToEmotes(emotes), nil
}

func (p *Provider) Channel(ctx context.Context, channelID string) ([]emote.Emote, error) {
	user, err := p.client.TwitchUser(ctx, channelID)
	if err != nil {
		return nil, fmt.Errorf("fetch twitch user: %w", err)
	}

	if user.EmoteSet == nil {
		return nil, nil
	}

	return p.emoteSetToEmotes(*user.EmoteSet), nil
}

func (p *Provider) emoteSetToEmotes(emoteSet seventv.EmoteSet) []emote.Emote {
	emotes := make([]emote.Emote, len(emoteSet.Emotes))

	for index, setEmote := range emoteSet.Emotes {
		emotes[index] = emote.Emote{
			ID:       setEmote.Id,
			Name:     setEmote.Name,
			Provider: "7tv",
		}
	}

	return emotes
}
