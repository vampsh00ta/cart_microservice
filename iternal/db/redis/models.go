package redis

type Cart map[string]MapItem

type Item struct {
	Id string `redis:"name"`
	MapItem
}

type MapItem struct {
	Name  string  `redis:"name"`
	Price float32 `redis:"price"`
	Image string  `redis:"image"`
}
