package github

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
	content.PublishTime = time.Now()
	log.Println("Publishe content to github")
	log.Println("Content: ", content.Body, content.Header, content.Footer)
	return nil
}
