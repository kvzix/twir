package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/twirapp/twir/libs/repositories/emotes"
)

type Repository struct {
	redis *redis.Client
}

var _ emotes.CacheRepository = (*Repository)(nil)

type Opts struct {
	Redis *redis.Client
}

func NewRepository(opts Opts) *Repository {
	return &Repository{
		redis: opts.Redis,
	}
}

func NewRepositoryFx(redis *redis.Client) *Repository {
	return NewRepository(Opts{
		Redis: redis,
	})
}

func (r *Repository) SetGlobalMany(
	ctx context.Context,
	emotes []emotes.SetEmoteInput,
	expiration time.Duration,
) error {
	_, err := r.redis.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, emote := range emotes {
			key := r.getGlobalEmoteKey(emote.Name)

			pipe.HSet(ctx, key, emoteModel{
				ID:       emote.ID,
				Name:     emote.Name,
				Provider: emote.Provider,
			})
			pipe.Expire(ctx, key, expiration)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("set emotes: %w", err)
	}

	return nil
}

func (r *Repository) SetChannelMany(
	ctx context.Context,
	channelID string,
	emotes []emotes.SetEmoteInput,
	expiration time.Duration,
) error {
	_, err := r.redis.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for _, emote := range emotes {
			key := r.getChannelEmoteKey(channelID, emote.Name)

			pipe.HSet(ctx, key, emoteModel{
				ID:       emote.ID,
				Name:     emote.Name,
				Provider: emote.Provider,
			})
			pipe.Expire(ctx, key, expiration)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("set emotes: %w", err)
	}

	return nil
}
