package blacklist

import (
	"io"

	"github.com/fvahid/sample-go/cache"
)

type Client interface {
	Load(blacklistFileName, cachePrefix string) error
	LoadFromReader(handle io.Reader, keyPrefix string) error
}
type Blacklist struct {
	CacheClient cache.Client
}

func (b *Blacklist) Load(blacklistFileName, cachePrefix string) error {
	return nil
}

func (b *Blacklist) LoadFromReader(handle io.Reader, keyPrefix string) error {
	return nil
}

func NewClient(cacheClient cache.Client) Client {
	return &Blacklist{CacheClient: cacheClient}
}
