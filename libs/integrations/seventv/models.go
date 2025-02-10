package seventv

type User struct {
	Id          string       `json:"id"`
	Username    string       `json:"username"`
	DisplayName string       `json:"display_name"`
	CreatedAt   int64        `json:"created_at"`
	AvatarUrl   string       `json:"avatar_url"`
	Style       struct{}     `json:"style"`
	Editors     []Editor     `json:"editors"`
	Roles       []string     `json:"roles"`
	Connections []Connection `json:"connections"`
}

type Editor struct {
	Id          string `json:"id"`
	Permissions int    `json:"permissions"`
	Visible     bool   `json:"visible"`
	AddedAt     int64  `json:"added_at"`
}

type Connection struct {
	Id            string      `json:"id"`
	Platform      string      `json:"platform"`
	Username      string      `json:"username"`
	DisplayName   string      `json:"display_name"`
	LinkedAt      int64       `json:"linked_at"`
	EmoteCapacity int         `json:"emote_capacity"`
	EmoteSetId    interface{} `json:"emote_set_id"`
	EmoteSet      *EmoteSet   `json:"emote_set"`
	User          User        `json:"user"`
}

type EmoteOwner struct {
	Id          string   `json:"id"`
	Username    string   `json:"username"`
	DisplayName string   `json:"display_name"`
	AvatarUrl   string   `json:"avatar_url"`
	Style       struct{} `json:"style"`
	Roles       []string `json:"roles"`
}

type EmoteSet struct {
	Id         string        `json:"id"`
	Name       string        `json:"name"`
	Flags      int           `json:"flags"`
	Tags       []interface{} `json:"tags"`
	Immutable  bool          `json:"immutable"`
	Privileged bool          `json:"privileged"`
	Capacity   int           `json:"capacity"`
	Owner      EmoteOwner    `json:"owner"`
	Emotes     []Emote       `json:"emotes"`
}

type Emote struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Flags     int    `json:"flags"`
	Timestamp int64  `json:"timestamp"`
	ActorId   string `json:"actor_id"`
	Data      struct {
		Id        string     `json:"id"`
		Name      string     `json:"name"`
		Flags     int        `json:"flags"`
		Tags      []string   `json:"tags"`
		Lifecycle int        `json:"lifecycle"`
		State     []string   `json:"state"`
		Listed    bool       `json:"listed"`
		Animated  bool       `json:"animated"`
		Owner     EmoteOwner `json:"owner"`
		Host      struct {
			Url   string `json:"url"`
			Files []struct {
				Name       string `json:"name"`
				StaticName string `json:"static_name"`
				Width      int    `json:"width"`
				Height     int    `json:"height"`
				FrameCount int    `json:"frame_count"`
				Size       int    `json:"size"`
				Format     string `json:"format"`
			} `json:"files"`
		} `json:"host"`
	} `json:"data"`
}
