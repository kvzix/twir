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

	activeChannels, err := s.channelsRepository.GetMany(
		ctx,
		channels.GetManyInput{
			Enabled: lo.ToPtr(true),
			Banned:  lo.ToPtr(false),
		},
	)
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
		return err
	}

	return nil
}

func (s *Service) SyncGlobalEmotes(
	ctx context.Context,
	expiration time.Duration,
) error {
	globalEmotes, err := s.getGlobalEmotes(ctx)
	if err != nil {
		return fmt.Errorf("get global emotes: %w", err)
	}

	err = s.emotesBatchLimiter.Batched(
		ctx,
		s.emotesToSetEmoteInputs(globalEmotes),
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

func (s *Service) SyncChannelEmotes(
	ctx context.Context,
	channelID string,
	expiration time.Duration,
) error {
	channelEmotes, err := s.getChannelEmotes(ctx, channelID)
	if err != nil {
		return fmt.Errorf("get channel emotes: %w", err)
	}

	err = s.emotesBatchLimiter.Batched(
		ctx,
		s.emotesToSetEmoteInputs(channelEmotes),
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
