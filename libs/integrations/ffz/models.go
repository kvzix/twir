package ffz

type ID = int64

type Emote struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
}

type EmoteSet struct {
	ID     ID      `json:"id"`
	Emotes []Emote `json:"emoticons"`
}

type EmoteSetCollection struct {
	EmoteSets map[string]EmoteSet `json:"sets"`
}
