package main

import (
	"fmt"
	"syscall"

	"github.com/fvahid/sample-go/blacklist"
	"github.com/fvahid/sample-go/cache"
	"github.com/go-redis/redis/v8"
)

var (
	redisAddress = envString("REDIS_ADDRESS", "localhost:6379")
)

func main() {
	ro := &redis.Options{
		Addr: redisAddress,
	}
	cacheClient := cache.NewClient(ro)
	blackListClient := blacklist.NewClient(cacheClient)
	fmt.Println(blackListClient)
}

func envString(key string, defaultVal string) string {
	if value, ok := syscall.Getenv(key); ok {
		return value
	}
	return defaultVal
}
