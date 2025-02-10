package emotescache

import (
	"github.com/satont/twir/apps/emotes-cacher/internal/emote"
	"github.com/twirapp/twir/libs/repositories/emotes"
)

func (s *Service) emotesToSetEmoteInputs(from []emote.Emote) []emotes.SetEmoteInput {
	inputs := make([]emotes.SetEmoteInput, len(from))

	for index, to := range from {
		inputs[index] = emotes.SetEmoteInput{
			ID:       to.ID,
			Name:     to.Name,
			Provider: to.Provider,
		}
	}

	return inputs
}
