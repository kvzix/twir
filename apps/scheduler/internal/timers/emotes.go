package timers

import (
	"context"
	"log/slog"
	"time"

	"github.com/satont/twir/libs/logger"
	buscore "github.com/twirapp/twir/libs/bus-core"
	"go.uber.org/fx"
)

const (
	week = 24 * 7 * time.Hour
)

type EmotesOpts struct {
	fx.In

	Lc     fx.Lifecycle
	Logger logger.Logger
	Bus    *buscore.Bus
}

func NewEmotes(opts EmotesOpts) {
	ticker := time.NewTicker(week)
	ctx, cancel := context.WithCancel(context.Background())

	opts.Lc.Append(
		fx.Hook{
			OnStart: func(_ context.Context) error {
				go func() {
					for {
						select {
						case <-ctx.Done():
							ticker.Stop()
							break
						case <-ticker.C:
							if err := opts.Bus.EmotesCacher.SyncEmotesCache.Publish(struct{}{}); err != nil {
								opts.Logger.Error("failed to sync emotes cache", slog.Any("error", err))
							}
						}
					}
				}()

				return nil
			},
			OnStop: func(_ context.Context) error {
				cancel()
				return nil
			},
		},
	)
}
