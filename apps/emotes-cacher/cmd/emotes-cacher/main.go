package main

import (
	"context"
	"log/slog"

	"github.com/satont/twir/apps/emotes-cacher/internal/delivery/bus"
	"github.com/satont/twir/apps/emotes-cacher/internal/emote"
	bttvprovider "github.com/satont/twir/apps/emotes-cacher/internal/emote/bttv"
	ffzprovider "github.com/satont/twir/apps/emotes-cacher/internal/emote/ffz"
	seventvprovider "github.com/satont/twir/apps/emotes-cacher/internal/emote/seventv"
	emotescache "github.com/satont/twir/apps/emotes-cacher/internal/services/emotes-cache"
	"github.com/satont/twir/libs/logger"
	"github.com/twirapp/twir/libs/baseapp"
	"github.com/twirapp/twir/libs/integrations/bttv"
	"github.com/twirapp/twir/libs/integrations/ffz"
	"github.com/twirapp/twir/libs/integrations/seventv"
	channelsrepository "github.com/twirapp/twir/libs/repositories/channels"
	channelsrepositorypgx "github.com/twirapp/twir/libs/repositories/channels/pgx"
	emotesrepository "github.com/twirapp/twir/libs/repositories/emotes"
	emotesrepositoryredis "github.com/twirapp/twir/libs/repositories/emotes/redis"
	"github.com/twirapp/twir/libs/uptrace"
	"go.uber.org/fx"
)

const service = "emotes-cacher"

func main() {
	fx.New(
		// Base
		baseapp.CreateBaseApp(
			baseapp.Opts{
				AppName: service,
			},
		),
		// Emotes
		fx.Provide(
			bttv.NewClient,
			ffz.NewClient,
			seventv.NewClient,
			fx.Annotate(
				bttvprovider.NewProvider,
				fx.As(new(emote.Provider)),
				fx.ResultTags(`group:"emote-providers"`),
			),
			fx.Annotate(
				ffzprovider.NewProvider,
				fx.As(new(emote.Provider)),
				fx.ResultTags(`group:"emote-providers"`),
			),
			fx.Annotate(
				seventvprovider.NewProvider,
				fx.As(new(emote.Provider)),
				fx.ResultTags(`group:"emote-providers"`),
			),
		),
		// Repositories
		fx.Provide(
			fx.Annotate(
				emotesrepositoryredis.NewRepositoryFx,
				fx.As(new(emotesrepository.CacheRepository)),
			),
			fx.Annotate(
				channelsrepositorypgx.NewFx,
				fx.As(new(channelsrepository.Repository)),
			),
		),
		// Services
		fx.Provide(
			emotescache.NewService,
		),
		// Runners
		fx.Invoke(
			uptrace.NewFx(service),
			bus.NewFx,
			func(logger logger.Logger, emotesCacheService *emotescache.Service) {
				if err := emotesCacheService.SyncEmotes(context.TODO()); err != nil {
					logger.Error("failed to sync emotes cache", slog.Any("error", err))
					return
				}
			},
		),
	).Run()
}
