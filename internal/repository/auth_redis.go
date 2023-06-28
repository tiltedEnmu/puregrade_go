package repository

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type AuthRedis struct {
	db *redis.Client
}

func NewAuthRedis(db *redis.Client) *AuthRedis {
	return &AuthRedis{db: db}
}

func (r *AuthRedis) UpsertRefreshToken(userId, token string) error {
	res := r.db.HSet(context.Background(), "refresh_tokens", userId, token)
	fmt.Print(res)
	return res.Err()
}

func (r *AuthRedis) GetRefreshToken(userId string) (string, error) {
	res := r.db.HGet(context.Background(), "refresh_tokens", userId)
	fmt.Print(res)
	return res.Result()
}
