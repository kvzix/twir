package bttv

type Emote struct {
	ID   string `json:"id"`
	Code string `json:"code"`
}

type User struct {
	ChannelEmotes []Emote `json:"channelEmotes"`
	SharedEmotes  []Emote `json:"sharedEmotes"`
}
