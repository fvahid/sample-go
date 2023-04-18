package redis

import (
	"log"
	"time"

	"github.com/fvahid/sample-go/config"
	"github.com/fvahid/sample-go/services"
)

type Publisher struct {
}

func NewPublisher(cfg config.Config) (services.Publisher, error) {
	return &Publisher{}, nil
}

func (p *Publisher) PublishContext(content services.Content) error {
	log.Println("Publishe content to file")
	content.PublishTime = time.Now()
	log.Println("Content: ", content)
	return nil
}
