package emotescache

import (
	"github.com/satont/twir/apps/emotes-cacher/internal/emote"
	"github.com/twirapp/twir/libs/repositories/emotes"
)

func (s *Service) emotesToSetEmoteInputs(inputEmotes []emote.Emote) []emotes.SetEmoteInput {
	inputs := make([]emotes.SetEmoteInput, len(inputEmotes))

	for index, inputEmote := range inputEmotes {
		inputs[index] = emotes.SetEmoteInput{
			ID:       inputEmote.ID,
			Name:     inputEmote.Name,
			Provider: inputEmote.Provider,
		}
	}

	return inputs
}
