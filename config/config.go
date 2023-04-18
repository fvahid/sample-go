package config

import "time"

type Config struct {
	CachePrefix   string
	RedisTimeout  time.Duration
	RedisHost     string
	RedisDB       int
	BlackListFile string
	FileName      string
	FilePath      string
	WatchPeriod   time.Duration
}
