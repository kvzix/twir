package emotescache

import (
	"context"
	"fmt"
	"sync"

	"github.com/satont/twir/apps/emotes-cacher/internal/emote"
	"golang.org/x/sync/errgroup"
)

type (
	EmotesFetchFunc func(emote.Provider) ([]emote.Emote, error)
)

func (s *Service) getGlobalEmotes(ctx context.Context) ([]emote.Emote, error) {
	return s.fetchEmotes(ctx, func(provider emote.Provider) ([]emote.Emote, error) {
		return provider.Global(ctx)
	})
}

func (s *Service) getChannelEmotes(ctx context.Context, channelID string) ([]emote.Emote, error) {
	return s.fetchEmotes(ctx, func(provider emote.Provider) ([]emote.Emote, error) {
		return provider.Channel(ctx, channelID)
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
			providerEmotes, err := fetch(emotesProvider)
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
		return nil, fmt.Errorf("fetchers: %w", err)
	}

	return emotes, nil
}
