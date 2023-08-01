package redis

import (
	"context"
	"github.com/google/uuid"
)

type Repository interface {
	AddItem(ctx context.Context, userId uuid.UUID, item Item) error
	DeleteItem(ctx context.Context, userId uuid.UUID, itemId string) error
	DeleteCart(ctx context.Context, userId uuid.UUID) error
	GetCart(ctx context.Context, userId uuid.UUID) (Cart, error)
}
