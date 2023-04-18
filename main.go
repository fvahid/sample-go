package main

import (
	"errors"
	"log"
	"syscall"

	"github.com/fvahid/sample-go/blacklist"
	"github.com/fvahid/sample-go/cache"
	"github.com/fvahid/sample-go/config"
	gitprovider "github.com/fvahid/sample-go/provider/github"
	redisprovider "github.com/fvahid/sample-go/provider/redis"
	gitpublisher "github.com/fvahid/sample-go/publisher/github"
	redispublisher "github.com/fvahid/sample-go/publisher/redis"
	"github.com/fvahid/sample-go/services"
	"github.com/go-redis/redis/v8"
)

var (
	redisAddress          = envString("REDIS_ADDRESS", "localhost:6379")
	ErrorIvalidProvider   = errors.New("invalid provider")
	ErrorInvalidPublisher = errors.New("invalid publisher")
)

func main() {
	var cfg config.Config
	ro := &redis.Options{
		Addr: redisAddress,
	}
	cacheClient := cache.NewClient(ro)
	blackListClient := blacklist.NewClient(cacheClient)
	err := blackListClient.Load(cfg.BlackListFile, cfg.CachePrefix)
	if err != nil {
		panic(err)
	}
	prov, err := getProvider("github")
	if err != nil {
		panic(err)
	}

	pubs, err := getPublisher("redis")
	if err != nil {
		panic(err)
	}
	s := services.Services{
		Publisher: pubs,
		Provider:  prov,
		Logger:    log.Default(),
	}
	s.Logger.Println("Running application logic")
	err = s.Run()
	if err != nil {
		panic(err)
	}

}

func getProvider(provider string) (services.Provider, error) {
	switch provider {
	case "github":
		{
			return gitprovider.NewProvider()
		}
	case "redis":
		{
			return redisprovider.NewProvider()
		}
	default:
		return nil, ErrorIvalidProvider
	}
}
func getPublisher(publisher string) (services.Publisher, error) {
	switch publisher {
	case "github":
		{
			return gitpublisher.NewPublisher()
		}
	case "redis":
		{
			return redispublisher.NewPublisher()
		}
	default:
		return nil, ErrorInvalidPublisher
	}
}

func envString(key string, defaultVal string) string {
	if value, ok := syscall.Getenv(key); ok {
		return value
	}
	return defaultVal
}
