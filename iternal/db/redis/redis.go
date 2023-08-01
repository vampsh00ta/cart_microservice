package redis

import (
	"github.com/go-kit/kit/log"
	//"github.com/go-kit/kit/log"
	"github.com/redis/go-redis/v9"
	//"github.com/go-kit/kit/log/level"
)

type Redis struct {
	client *redis.Client
	logger log.Logger
}

func New(client *redis.Client, logger log.Logger) Repository {
	return &Redis{
		client: client,
		logger: logger,
	}
}
