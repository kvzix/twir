package emotescache

import (
	"context"
	"fmt"
	"sync"

	"github.com/satont/twir/apps/emotes-cacher/internal/emote"
	"golang.org/x/sync/errgroup"
)

type (
	EmotesFetchFunc func(context.Context, emote.Provider) ([]emote.Emote, error)
)

func (s *Service) fetchGlobalEmotes(ctx context.Context) ([]emote.Emote, error) {
	return s.fetchEmotes(ctx, func(ctx context.Context, provider emote.Provider) ([]emote.Emote, error) {
		globalEmotes, err := provider.GlobalEmotes(ctx)
		if err != nil {
			return nil, err
		}

		return globalEmotes, nil
	})
}

func (s *Service) fetchChannelEmotes(ctx context.Context, channelID string) ([]emote.Emote, error) {
	return s.fetchEmotes(ctx, func(ctx context.Context, provider emote.Provider) ([]emote.Emote, error) {
		channelEmotes, err := provider.ChannelEmotes(ctx, channelID)
		if err != nil {
			return nil, err
		}

		return channelEmotes, nil
	})
}

func (s *Service) fetchEmotes(ctx context.Context, fetch EmotesFetchFunc) ([]emote.Emote, error) {
	fetchers, ctx := errgroup.WithContext(ctx)

	var (
		emotes       = make([]emote.Emote, 0, 100)
		emotesLocker sync.Mutex
	)

	for _, emotesProvider := range s.emotesProviders {
		fetchers.Go(func() error {
			providerEmotes, err := fetch(ctx, emotesProvider)
			if err != nil {
				return fmt.Errorf("fetch emotes from provider: %w", err)
			}

			emotesLocker.Lock()
			emotes = append(emotes, providerEmotes...)
			emotesLocker.Unlock()

			return nil
		})
	}

	if err := fetchers.Wait(); err != nil {
		return nil, err
	}

	return emotes, nil
}
