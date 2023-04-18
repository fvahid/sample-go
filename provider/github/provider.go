package github

import (
	"log"
	"time"

	"github.com/fvahid/sample-go/config"
	"github.com/fvahid/sample-go/services"
)

type Provider struct {
}

func NewProvider(cfg config.Config) (services.Provider, error) {
	return &Provider{}, nil
}
func (p *Provider) ProvideContent() (services.Content, error) {
	log.Println("Provide content from github")
	return services.Content{
		Body:        "",
		Header:      "",
		Footer:      "",
		ProvideTime: time.Now(),
	}, nil
}
