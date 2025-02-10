package app

import (
	bus_listener "github.com/satont/twir/apps/scheduler/internal/bus-listener"
	"github.com/satont/twir/apps/scheduler/internal/services"
	"github.com/satont/twir/apps/scheduler/internal/timers"
	config "github.com/satont/twir/libs/config"
	"github.com/satont/twir/libs/logger"
	"github.com/twirapp/twir/libs/baseapp"
	"github.com/twirapp/twir/libs/grpc/clients"
	"github.com/twirapp/twir/libs/grpc/parser"
	"github.com/twirapp/twir/libs/grpc/tokens"
	"github.com/twirapp/twir/libs/uptrace"
	"go.uber.org/fx"
)

const service = "scheduler"

var App = fx.Module(
	service,
	baseapp.CreateBaseApp(baseapp.Opts{AppName: service}),
	fx.Provide(
		func(c config.Config) parser.ParserClient {
			return clients.NewParser(c.AppEnv)
		},
		func(c config.Config) tokens.TokensClient {
			return clients.NewTokens(c.AppEnv)
		},
		services.NewRoles,
		services.NewCommands,
	),
	fx.Invoke(
		uptrace.NewFx(service),
		bus_listener.New,
		timers.NewOnlineUsers,
		timers.NewStreams,
		timers.NewCommandsAndRoles,
		timers.NewBannedChannels,
		timers.NewWatched,
		timers.NewExpiredCommands,
		func(l logger.Logger) {
			l.Info("Started")
		},
	),
)
