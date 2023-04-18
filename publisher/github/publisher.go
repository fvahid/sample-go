package github

import (
	"log"

	"github.com/fvahid/sample-go/services"
)

type Publisher struct {
}

func NewPublisher() (services.Publisher, error) {
	return &Publisher{}, nil
}

func (p *Publisher) PublishContext(content services.Content) error {
	log.Println("Publishe content to github")
	log.Println("Content: ", content.Body, content.Header, content.Footer)
	return nil
}
