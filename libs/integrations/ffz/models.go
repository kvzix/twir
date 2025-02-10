package ffz

type Emote struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type EmoteSet struct {
	ID     string  `json:"id"`
	Emotes []Emote `json:"emoticons"`
}

type EmoteSetCollection struct {
	EmoteSets map[string]EmoteSet `json:"sets"`
}
