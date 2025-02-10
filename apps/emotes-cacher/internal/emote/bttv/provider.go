package bttv

import (
	"context"
	"fmt"

	"github.com/satont/twir/apps/emotes-cacher/internal/emote"
	"github.com/twirapp/twir/libs/integrations/bttv"
)

type Provider struct {
	client bttv.Client
}

var _ emote.Provider = (*Provider)(nil)

func NewProvider(client bttv.Client) *Provider {
	return &Provider{
		client: client,
	}
}

func (p *Provider) Global(ctx context.Context) ([]emote.Emote, error) {
	globalEmotes, err := p.client.GlobalEmotes(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetch global emotes: %w", err)
	}

	emotes := make([]emote.Emote, len(globalEmotes))

	for index, globalEmote := range globalEmotes {
		emotes[index] = p.emoteToEntity(globalEmote)
	}

	return emotes, nil
}

func (p *Provider) Channel(ctx context.Context, channelID string) ([]emote.Emote, error) {
	user, err := p.client.TwitchUser(ctx, channelID)
	if err != nil {
		return nil, fmt.Errorf("fetch twitch user: %w", err)
	}

	emotes := make([]emote.Emote, 0, len(user.ChannelEmotes)+len(user.SharedEmotes))

	for _, channelEmote := range user.ChannelEmotes {
		emotes = append(emotes, p.emoteToEntity(channelEmote))
	}

	for _, sharedEmote := range user.SharedEmotes {
		emotes = append(emotes, p.emoteToEntity(sharedEmote))
	}

	return emotes, nil
}

func (p *Provider) emoteToEntity(from bttv.Emote) emote.Emote {
	return emote.Emote{
		ID:       from.ID,
		Name:     from.Code,
		Provider: "bttv",
	}
}
