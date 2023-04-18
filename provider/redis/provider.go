package redis

import (
	"log"

	"github.com/fvahid/sample-go/services"
)

type Provider struct {
}

func NewProvider() (services.Provider, error) {
	return &Provider{}, nil
}
func (p *Provider) ProvideContent() (services.Content, error) {
	log.Println("Provide content from redis")
	return services.Content{
		Body:   "",
		Header: "",
		Footer: "",
	}, nil
}
