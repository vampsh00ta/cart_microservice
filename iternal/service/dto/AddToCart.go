package dto

import (
	"cart_mircoservice/iternal/db/redis"
	"github.com/google/uuid"
)

type AddToCartRequest struct {
	UserId uuid.UUID `json:"userId,omitempty" validate:"required"`
	redis.Item
}

type AddToCartResponse struct {
	UserId uuid.UUID  `json:"userId,omitempty"`
	Item   redis.Item `json:"item,omitempty"`
	Err    error      `json:"error,omitempty"`
}

type GRPCAddToCartRequest struct {
	JWT  string `json:"jwt,omitempty" validate:"required,string"`
	Item redis.Item
}

type GRPCAddToCartResponse struct {
	UserId uuid.UUID  `json:"userId,omitempty"`
	Item   redis.Item `json:"item,omitempty"`
}
