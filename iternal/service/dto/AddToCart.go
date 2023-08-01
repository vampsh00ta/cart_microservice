package dto

import (
	"cart_mircoservice/iternal/db/redis"
	"github.com/google/uuid"
)

type AddToCartRequest struct {
	UserId uuid.UUID `json:"userId"`
	redis.Item
}

type AddToCartResponse struct {
	UserId uuid.UUID  `json:"userId,omitempty"`
	Item   redis.Item `json:"item,omitempty"`
	Err    error      `json:"error,omitempty"`
}
