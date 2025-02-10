package emotescache

import (
	"context"
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/twirapp/twir/libs/repositories/channels"
	"github.com/twirapp/twir/libs/repositories/emotes"
	"golang.org/x/sync/errgroup"
)

func (s *Service) SyncEmotes(ctx context.Context) error {
	if err := s.SyncGlobalEmotes(ctx, s.config.EmotesCacherEmoteTTL); err != nil {
		return fmt.Errorf("sync global emotes: %w", err)
	}

	input := channels.GetManyInput{
		Enabled: lo.ToPtr(true),
		Banned:  lo.ToPtr(false),
	}

	activeChannels, err := s.channelsRepository.GetMany(ctx, input)
	if err != nil {
		return fmt.Errorf("get active channels: %w", err)
	}

	syncers, ctx := errgroup.WithContext(ctx)

	for _, activeChannel := range activeChannels {
		syncers.Go(func() error {
			if err = s.SyncChannelEmotes(
				ctx,
				activeChannel.ID,
				s.config.EmotesCacherEmoteTTL,
			); err != nil {
				return fmt.Errorf("sync channel emotes: %w", err)
			}

			return nil
		})
	}

	if err = syncers.Wait(); err != nil {
		return fmt.Errorf("syncers: %w", err)
	}

	return nil
}

func (s *Service) SyncGlobalEmotes(ctx context.Context, expiration time.Duration) error {
	globalEmotes, err := s.fetchGlobalEmotes(ctx)
	if err != nil {
		return fmt.Errorf("fetch global emotes: %w", err)
	}

	inputs := s.emotesToSetEmoteInputs(globalEmotes)

	err = s.emotesBatchLimiter.Batched(
		ctx,
		inputs,
		func(ctx context.Context, batch []emotes.SetEmoteInput) error {
			if err = s.emotesCacheRepository.SetGlobalMany(ctx, batch, expiration); err != nil {
				return fmt.Errorf("set global emotes: %w", err)
			}

			return nil
		},
	)
	if err != nil {
		return fmt.Errorf("batch: %w", err)
	}

	return nil
}

func (s *Service) SyncChannelEmotes(ctx context.Context, channelID string, expiration time.Duration) error {
	channelEmotes, err := s.fetchChannelEmotes(ctx, channelID)
	if err != nil {
		return fmt.Errorf("fetch channel emotes: %w", err)
	}

	inputs := s.emotesToSetEmoteInputs(channelEmotes)

	err = s.emotesBatchLimiter.Batched(
		ctx,
		inputs,
		func(ctx context.Context, batch []emotes.SetEmoteInput) error {
			if err = s.emotesCacheRepository.SetChannelMany(ctx, channelID, batch, expiration); err != nil {
				return fmt.Errorf("set channel emotes: %w", err)
			}

			return nil
		},
	)
	if err != nil {
		return fmt.Errorf("batch: %w", err)
	}

	return nil
}
