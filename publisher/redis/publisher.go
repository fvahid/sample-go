package redis

import (
	"context"
	"log"
	"time"

	"github.com/fvahid/sample-go/config"
	"github.com/fvahid/sample-go/services"
	"github.com/go-redis/redis/v8"
)

type Publisher struct {
	RedisHost    string
	RedisTimeout time.Duration
	RedisDB      int
	CachePrefix  string
}

func NewPublisher(cfg config.Config) (services.Publisher, error) {
	return &Publisher{
		CachePrefix:  cfg.CachePrefix,
		RedisHost:    cfg.RedisHost,
		RedisTimeout: cfg.RedisTimeout,
		RedisDB:      cfg.RedisDB,
	}, nil
}

func (p *Publisher) PublishContext(content services.Content) error {
	content.PublishTime = time.Now()
	log.Println("Publishe content to redis")
	return p.redisSave(content.Body)
}

func (p *Publisher) redisSave(msg string) error {
	r := redis.NewClient(&redis.Options{Addr: p.RedisHost, DB: p.RedisDB})
	return r.Set(context.Background(), p.CachePrefix, msg, p.RedisTimeout).Err()
}
