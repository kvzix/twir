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
	emotes, err := p.client.GlobalEmotes(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetch global emotes: %w", err)
	}

	emotesIDs := make([]emote.Emote, len(emotes))

	for index, globalEmote := range emotes {
		emotesIDs[index] = p.fromEmote(globalEmote)
	}

	return emotesIDs, nil
}

func (p *Provider) Channel(ctx context.Context, channelID string) ([]emote.Emote, error) {
	user, err := p.client.TwitchUser(ctx, channelID)
	if err != nil {
		return nil, fmt.Errorf("fetch twitch user: %w", err)
	}

	emotesIDs := make([]emote.Emote, 0, len(user.ChannelEmotes)+len(user.SharedEmotes))

	for _, channelEmote := range user.ChannelEmotes {
		emotesIDs = append(emotesIDs, p.fromEmote(channelEmote))
	}

	for _, sharedEmote := range user.SharedEmotes {
		emotesIDs = append(emotesIDs, p.fromEmote(sharedEmote))
	}

	return emotesIDs, nil
}

func (p *Provider) fromEmote(from bttv.Emote) emote.Emote {
	return emote.Emote{
		ID:       from.ID,
		Name:     from.Code,
		Provider: "bttv",
	}
}
