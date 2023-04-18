package services

import (
	"log"
	"time"
)

type Content struct {
	Body        string
	Header      string
	Footer      string
	ProvideTime time.Time
	PublishTime time.Time
}
type Provider interface {
	ProvideContent() (Content, error)
}
type Publisher interface {
	PublishContext(content Content) error
}
type Services struct {
	Publisher Publisher
	Provider  Provider
	Logger    *log.Logger
}

func (s Services) Run() error {
	msg, err := s.Provider.ProvideContent()
	if err != nil {
		s.Logger.Println("can't provide content")
		return err
	}
	err = s.Publisher.PublishContext(msg)
	if err != nil {
		s.Logger.Println("can't publsh content")
		return err
	}
	return nil
}
