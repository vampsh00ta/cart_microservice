package dto

import (
	"github.com/google/uuid"

	"cart_mircoservice/iternal/db/redis"
)

type GetFromCartRequest struct {
	UserId uuid.UUID `json:"userId"`
}

type GetFromCartResponse struct {
	Cart redis.Cart `json:"cart,omitempty"`
	Err  error      `json:"error,omitempty"`
}
