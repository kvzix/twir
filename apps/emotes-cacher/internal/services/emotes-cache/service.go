package emotescache

import (
	"github.com/satont/twir/apps/emotes-cacher/internal/emote"
	limitedbatcher "github.com/satont/twir/apps/emotes-cacher/pkg/limited-batcher"
	config "github.com/satont/twir/libs/config"
	"github.com/twirapp/twir/libs/repositories/channels"
	"github.com/twirapp/twir/libs/repositories/emotes"
	"go.uber.org/fx"
)

type Service struct {
	config config.Config

	emotesProviders       []emote.Provider
	emotesLimitedBatcher  limitedbatcher.LimitedBatcher[emotes.SetEmoteInput]
	emotesCacheRepository emotes.CacheRepository

	channelsRepository channels.Repository
}

type Params struct {
	fx.In

	Config                config.Config
	EmotesProviders       []emote.Provider `group:"emote-providers"`
	EmotesCacheRepository emotes.CacheRepository
	ChannelsRepository    channels.Repository
}

func NewService(params Params) *Service {
	emotesLimitedBatcher := limitedbatcher.New[emotes.SetEmoteInput](
		params.Config.EmotesCacherBatchRate,
		params.Config.EmotesCacherBatchSize,
	)

	return &Service{
		config:                params.Config,
		emotesProviders:       params.EmotesProviders,
		emotesLimitedBatcher:  emotesLimitedBatcher,
		emotesCacheRepository: params.EmotesCacheRepository,
		channelsRepository:    params.ChannelsRepository,
	}
}
