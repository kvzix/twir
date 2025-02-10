package emotescache

import (
	batchlimiter "github.com/satont/twir/apps/emotes-cacher/internal/batch-limiter"
	"github.com/satont/twir/apps/emotes-cacher/internal/emote"
	config "github.com/satont/twir/libs/config"
	"github.com/twirapp/twir/libs/repositories/channels"
	"github.com/twirapp/twir/libs/repositories/emotes"
	"go.uber.org/fx"
)

type Service struct {
	config config.Config

	emotesProviders       []emote.Provider
	emotesBatchLimiter    batchlimiter.BatchLimiter[emotes.SetEmoteInput]
	emotesCacheRepository emotes.CacheRepository

	channelsRepository channels.Repository
}

type Params struct {
	fx.In

	Config                config.Config
	EmoteProviders        []emote.Provider `group:"emote-providers"`
	EmotesCacheRepository emotes.CacheRepository
	ChannelsRepository    channels.Repository
}

func NewService(params Params) *Service {
	emotesBatchLimiter := batchlimiter.New[emotes.SetEmoteInput](
		params.Config.EmotesCacherBatchRate,
		params.Config.EmotesCacherBatchSize,
	)

	return &Service{
		config:                params.Config,
		emotesProviders:       params.EmoteProviders,
		emotesBatchLimiter:    emotesBatchLimiter,
		emotesCacheRepository: params.EmotesCacheRepository,
		channelsRepository:    params.ChannelsRepository,
	}
}
