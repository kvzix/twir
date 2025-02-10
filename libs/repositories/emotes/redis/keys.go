package redis

import (
	"fmt"
)

func (r *Repository) getGlobalEmoteKey(emote string) string {
	return fmt.Sprintf("emotes:global:%s", emote)
}

func (r *Repository) getChannelEmoteKey(channelID, emote string) string {
	return fmt.Sprintf("emotes:channel:%s:%s", channelID, emote)
}
