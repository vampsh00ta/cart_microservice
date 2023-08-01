package service

import (
	"cart_mircoservice/iternal/db/redis"
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/google/uuid"
)

type service struct {
	db     redis.Repository
	logger log.Logger
}

func (s *service) AddToCart(ctx context.Context, userId uuid.UUID, item redis.Item) error {
	logger := log.With(s.logger, "method", "Create")

	err := s.db.AddItem(ctx, userId, item)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	return nil
}

func (s *service) DeleteFromCart(ctx context.Context, userId uuid.UUID, itemId string) error {
	logger := log.With(s.logger, "method", "DeleteFromCart")
	err := s.db.DeleteItem(ctx, userId, itemId)
	if err != nil {
		level.Error(logger).Log("err", err)
		return err
	}
	return err
}

func (s *service) GetFromCart(ctx context.Context, userId uuid.UUID) (redis.Cart, error) {
	logger := log.With(s.logger, "method", "GetFromCart")
	cart, err := s.db.GetCart(ctx, userId)
	if err != nil {
		level.Error(logger).Log("err", err)
		return redis.Cart{}, nil
	}
	return cart, nil
}

func New(db redis.Repository, logger log.Logger) Service {
	return &service{
		db:     db,
		logger: logger,
	}
}
