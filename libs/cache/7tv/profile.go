package seventv

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	generic_cacher "github.com/twirapp/twir/libs/cache/generic-cacher"
	seventvintegration "github.com/twirapp/twir/libs/integrations/seventv"
)

func New(
	redis *redis.Client,
) *generic_cacher.GenericCacher[*seventvintegration.Connection] {
	return generic_cacher.New[*seventvintegration.Connection](
		generic_cacher.Opts[*seventvintegration.Connection]{
			Redis:     redis,
			KeyPrefix: "cache:twir:seventv:profile:",
			LoadFn: func(ctx context.Context, key string) (
				*seventvintegration.Connection,
				error,
			) {
				profile, err := seventvintegration.GetProfile(ctx, key)
				if err != nil {
					return nil, err
				}

				return &profile, nil
			},
			Ttl: 5 * time.Minute,
		},
	)
}

func NewBySeventvID(
	redis *redis.Client,
) *generic_cacher.GenericCacher[*seventvintegration.Connection] {
	return generic_cacher.New[*seventvintegration.Connection](
		generic_cacher.Opts[*seventvintegration.Connection]{
			Redis:     redis,
			KeyPrefix: "cache:twir:seventv:profile:by-seventv-id:",
			LoadFn: func(ctx context.Context, key string) (
				*seventvintegration.Connection,
				error,
			) {
				profile, err := seventvintegration.GetProfileBySevenTvID(ctx, key)
				if err != nil {
					return nil, err
				}

				return &profile, nil
			},
			Ttl: 5 * time.Minute,
		},
	)
}
