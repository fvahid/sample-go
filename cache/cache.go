package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type Client interface {
	Set(key string, value interface{}, exp time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
	Scan(key string, action func(context.Context, string) error) error
}

type Repository struct {
	Client *redis.Client
}

func (r *Repository) Set(key string, value interface{}, exp time.Duration) error {
	return nil
}
func (r *Repository) Get(key string) (string, error) {
	return "", nil
}
func (r *Repository) Del(key string) error {
	return nil
}
func (r *Repository) Scan(key string, action func(context.Context, string) error) error {
	return nil
}

func NewClient(ro *redis.Options) Client {
	return &Repository{
		Client: redis.NewClient(ro),
	}
}
