package main

import (
	"errors"
	"log"
	"strconv"
	"syscall"
	"time"

	"github.com/fvahid/sample-go/blacklist"
	"github.com/fvahid/sample-go/cache"
	"github.com/fvahid/sample-go/config"
	fileprovider "github.com/fvahid/sample-go/provider/file"
	gitprovider "github.com/fvahid/sample-go/provider/github"
	redisprovider "github.com/fvahid/sample-go/provider/redis"
	filepublisher "github.com/fvahid/sample-go/publisher/file"
	gitpublisher "github.com/fvahid/sample-go/publisher/github"
	redispublisher "github.com/fvahid/sample-go/publisher/redis"

	"github.com/fvahid/sample-go/services"
	"github.com/go-redis/redis/v8"
)

var (
	redisAddress = envString("REDIS_ADDRESS", "localhost:6379")
	redisTimeout = envString("REDIS_TIMEOUT", "10s")
	redisDB      = envString("REDIS_DB", "0")
	fileName     = envString("FILE_NAME", "main.go")
	filePath     = envString("PWD", "")
	cachePrefix  = envString("CACHE_PREFIX", "CachePrefix")
	watchPeriod  = envString("WATCH_PERIOD", "10s")

	ErrorIvalidProvider   = errors.New("invalid provider")
	ErrorInvalidPublisher = errors.New("invalid publisher")
)

func initConfig(cfg config.Config) config.Config {
	cfg.FileName = fileName
	cfg.FilePath = filePath
	cfg.CachePrefix = cachePrefix
	timeout, err := time.ParseDuration(redisTimeout)
	if err != nil {
		panic(err)
	}
	cfg.RedisTimeout = timeout
	cfg.RedisHost = redisAddress
	redisDatabase, err := strconv.Atoi(redisDB)
	if err != nil {
		panic(err)
	}
	cfg.RedisDB = redisDatabase

	watch, err := time.ParseDuration(watchPeriod)
	if err != nil {
		panic(err)
	}
	cfg.WatchPeriod = watch
	return cfg
}

func main() {
	var initconfig config.Config
	cfg := initConfig(initconfig)

	ro := &redis.Options{
		Addr: redisAddress,
	}
	cacheClient := cache.NewClient(ro)
	blackListClient := blacklist.NewClient(cacheClient)
	err := blackListClient.Load(cfg.BlackListFile, cfg.CachePrefix)
	if err != nil {
		panic(err)
	}

	prov, err := getProvider("file", cfg)
	if err != nil {
		panic(err)
	}

	pubs, err := getPublisher("redis", cfg)
	if err != nil {
		panic(err)
	}
	s := services.Services{
		Publisher: pubs,
		Provider:  prov,
		Logger:    log.Default(),
	}
	s.Logger.Println("Running application logic")
	for {
		err = s.Run()
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Duration(cfg.WatchPeriod))
	}

}

func getProvider(provider string, cfg config.Config) (services.Provider, error) {
	switch provider {
	case "github":
		return gitprovider.NewProvider(cfg)
	case "redis":
		return redisprovider.NewProvider(cfg)
	case "file":
		return fileprovider.NewProvider(cfg)
	default:
		return nil, ErrorIvalidProvider
	}
}
func getPublisher(publisher string, cfg config.Config) (services.Publisher, error) {
	switch publisher {
	case "github":
		return gitpublisher.NewPublisher(cfg)
	case "redis":
		return redispublisher.NewPublisher(cfg)
	case "file":
		return filepublisher.NewPublisher(cfg)
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
