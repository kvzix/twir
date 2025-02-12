package bus

import (
	"context"
	"fmt"

	"go.uber.org/fx"
)

func NewFx(params Params, lc fx.Lifecycle) Bus {
	bus := New(params)

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				if err := bus.bus.EmotesCacher.SyncEmotesCache.SubscribeGroup(
					"emotes-cacher",
					bus.syncEmotesCache,
				); err != nil {
					return fmt.Errorf("subscribe to sync emotes cache: %w", err)
				}

				return nil
			},
			OnStop: func(ctx context.Context) error {
				bus.bus.EmotesCacher.SyncEmotesCache.Unsubscribe()
				return nil
			},
		},
	)

	return bus
}
