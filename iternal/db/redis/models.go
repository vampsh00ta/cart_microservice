package redis

type Cart map[string]MapItem

type Item struct {
	Id string `redis:"id" validate:"required"`
	MapItem
}

type MapItem struct {
	Name  string  `redis:"name" validate:"required"`
	Price float32 `redis:"price" validate:"required"`
	Image string  `redis:"image" validate:"required"`
}
