package bus

import (
	"context"
	"log/slog"

	emotescache "github.com/satont/twir/apps/emotes-cacher/internal/services/emotes-cache"
	"github.com/satont/twir/libs/logger"
	buscore "github.com/twirapp/twir/libs/bus-core"
	"go.uber.org/fx"
)

type Bus struct {
	bus                *buscore.Bus
	logger             logger.Logger
	emotesCacheService *emotescache.Service
}

type Params struct {
	fx.In

	Bus    *buscore.Bus
	Logger logger.Logger

	EmotesCacheService *emotescache.Service
}

func New(params Params) Bus {
	return Bus{
		bus:                params.Bus,
		logger:             params.Logger,
		emotesCacheService: params.EmotesCacheService,
	}
}

func (b *Bus) syncEmotesCache(ctx context.Context, _ struct{}) struct{} {
	if err := b.emotesCacheService.SyncEmotes(ctx); err != nil {
		b.logger.Error("failed to sync emotes cache", slog.Any("error", err))
	}

	return struct{}{}
}
