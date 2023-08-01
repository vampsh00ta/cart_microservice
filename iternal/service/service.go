package service

import (
	db "cart_mircoservice/iternal/db/redis"
	"context"
	"github.com/google/uuid"
)

type Service interface {
	AddToCart(ctx context.Context, userId uuid.UUID, item db.Item) error
	DeleteFromCart(ctx context.Context, userId uuid.UUID, itemId string) error
	GetFromCart(ctx context.Context, userId uuid.UUID) (db.Cart, error)
}
