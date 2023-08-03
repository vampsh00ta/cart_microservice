package redis

import (
	"cart_mircoservice/iternal/config"
	"context"
	"encoding/json"
	"github.com/google/uuid"
)

func (r *Redis) DeleteCart(ctx context.Context, userId uuid.UUID) error {
	if err := r.client.Del(ctx, "user_"+userId.String()).Err(); err != nil {
		return err
	}
	return nil
}
func (r *Redis) AddItem(ctx context.Context, userId uuid.UUID, item Item) error {
	cart, err := r.GetCart(ctx, userId)
	if err != nil {
		return err
	}
	if _, ok := cart[item.Id]; !ok {
		cart[item.Id] = item.MapItem
	} else {
		return config.AlreadyInCart
	}
	cartToSet, err := json.Marshal(&cart)
	if err != nil {
		return err
	}
	if err := r.client.Set(ctx, "user_"+userId.String(), string(cartToSet), 0).Err(); err != nil {
		return err
	}
	return nil

}
func (r *Redis) GetCart(ctx context.Context, userId uuid.UUID) (Cart, error) {
	cartByte, err := r.client.Get(ctx, "user_"+userId.String()).Bytes()
	if cartByte == nil {
		return Cart{}, nil
	}
	if err != nil {
		return nil, err
	}
	var cart Cart
	json.Unmarshal(cartByte, &cart)

	return cart, nil

}

func (r *Redis) DeleteItem(ctx context.Context, userId uuid.UUID, itemId string) error {
	cart, err := r.GetCart(ctx, userId)
	if err != nil {
		return err
	}
	delete(cart, itemId)
	cartToSet, err := json.Marshal(&cart)
	if err != nil {
		return err
	}
	if err := r.client.Set(ctx, "user_"+userId.String(), string(cartToSet), 0).Err(); err != nil {
		return err
	}
	return nil
}
