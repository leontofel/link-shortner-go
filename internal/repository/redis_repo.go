package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisRepository struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisRepository(ctx context.Context, addr string) *RedisRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisRepository{client: rdb, ctx: ctx}
}

func (r *RedisRepository) Save(code string, url string) error {
	return r.client.Set(r.ctx, code, url, 24*time.Hour).Err()
}

func (r *RedisRepository) Find(code string) (string, error) {
	return r.client.Get(r.ctx, code).Result()
}
