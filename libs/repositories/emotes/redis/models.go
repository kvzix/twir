package redis

type emoteModel struct {
	ID       string `redis:"id"`
	Name     string `redis:"name"`
	Provider string `redis:"provider"`
}
