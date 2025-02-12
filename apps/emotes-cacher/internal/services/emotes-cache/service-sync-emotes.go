package emotescache

import (
	"context"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v5"
	"github.com/samber/lo"
	"github.com/twirapp/twir/libs/repositories/channels"
	"github.com/twirapp/twir/libs/repositories/emotes"
	"golang.org/x/sync/errgroup"
)

var (
	syncEmotesMaxTries uint = 15

	syncEmotesBackOff = &backoff.ExponentialBackOff{
		InitialInterval:     1 * time.Second,
		RandomizationFactor: 0.5,
		Multiplier:          1.5,
		MaxInterval:         60 * time.Second,
	}
)

func (s *Service) SyncEmotes(ctx context.Context) error {
	syncers, _ := errgroup.WithContext(ctx)

	syncers.Go(
		func() error {
			return s.retrySyncEmotes(ctx, func() error {
				return s.SyncGlobalEmotes(ctx, s.config.EmotesCacherEmoteTTL)
			})
		},
	)

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

	for _, channel := range activeChannels {
		syncers.Go(
			func() error {
				return s.retrySyncEmotes(ctx, func() error {
					return s.SyncChannelEmotes(ctx, channel.ID, s.config.EmotesCacherEmoteTTL)
				})
			},
		)
	}

	return syncers.Wait()
}

func (s *Service) SyncGlobalEmotes(ctx context.Context, expiration time.Duration) error {
	globalEmotes, err := s.fetchGlobalEmotes(ctx)
	if err != nil {
		return fmt.Errorf("fetch global emotes: %w", err)
	}

	inputs := s.emotesToSetEmoteInputs(globalEmotes)

	return s.emotesLimitedBatcher.Batch(
		ctx,
		inputs,
		func(ctx context.Context, batch []emotes.SetEmoteInput) error {
			if err = s.emotesCacheRepository.SetGlobalMany(ctx, batch, expiration); err != nil {
				return fmt.Errorf("set global emotes: %w", err)
			}

			return nil
		},
	)
}

func (s *Service) SyncChannelEmotes(ctx context.Context, channelID string, expiration time.Duration) error {
	channelEmotes, err := s.fetchChannelEmotes(ctx, channelID)
	if err != nil {
		return fmt.Errorf("fetch channel emotes: %w", err)
	}

	inputs := s.emotesToSetEmoteInputs(channelEmotes)

	return s.emotesLimitedBatcher.Batch(
		ctx,
		inputs,
		func(ctx context.Context, batch []emotes.SetEmoteInput) error {
			if err = s.emotesCacheRepository.SetChannelMany(ctx, channelID, batch, expiration); err != nil {
				return fmt.Errorf("set channel emotes: %w", err)
			}

			return nil
		},
	)
}

func (s *Service) retrySyncEmotes(ctx context.Context, sync func() error) error {
	if _, err := backoff.Retry(
		ctx,
		func() (struct{}, error) {
			return struct{}{}, sync()
		},
		backoff.WithMaxTries(syncEmotesMaxTries),
		backoff.WithBackOff(syncEmotesBackOff),
	); err != nil {
		return err
	}

	return nil
}
