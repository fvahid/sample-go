package redis

import (
	"io/ioutil"
	"log"
	"time"

	"github.com/fvahid/sample-go/config"
	"github.com/fvahid/sample-go/services"
)

type Provider struct {
	FileName string
	FilePath string
}

func NewProvider(cfg config.Config) (services.Provider, error) {
	return &Provider{FileName: cfg.FileName, FilePath: cfg.FilePath}, nil
}
func (p *Provider) ProvideContent() (services.Content, error) {
	log.Println("Provide content from file")
	fileContent, err := p.readFileContent()
	if err != nil {
		return services.Content{}, err
	}
	return services.Content{
		Body:        fileContent,
		Header:      "",
		Footer:      "",
		ProvideTime: time.Now(),
	}, nil
}
func (p *Provider) readFileContent() (string, error) {
	content, err := ioutil.ReadFile(p.FilePath + "/" + p.FileName)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
