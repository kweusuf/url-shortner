package url

import (
	"context"
	"fmt"

	client "github.com/kweusuf/url-shortner/pkg/client/url"
	"github.com/kweusuf/url-shortner/pkg/utils/log"
)

type URLService interface {
	ShortenURLService(context context.Context, url string) (string, error)
	ExpandURLService(context context.Context, url string) (string, error)
}

type urlService struct {
	client client.URLClient
}

func MakeURLService(client client.URLClient) URLService {
	return &urlService{
		client: client,
	}
}

// ShortenURLService implements URLService.
func (s *urlService) ShortenURLService(context context.Context, url string) (string, error) {
	log.Debug(fmt.Sprintf("In ShortenURLService with URL: %s", url))
	short, err := s.client.ShortenURL(context, url)
	if err != nil {
		log.Error("Error shortening URL: ", err)
		return "", err
	}
	log.Debug(fmt.Sprintf("Shortened URL response: %v", short))
	return short.(string), nil
}

// ExpandURLService implements URLService.
func (s *urlService) ExpandURLService(context context.Context, url string) (string, error) {
	log.Debug(fmt.Sprintf("In ExpandURLService with URL: %s", url))
	expanded, err := s.client.ExpandURL(context, url)
	if err != nil {
		log.Error("Error expanding URL: ", err)
		return "", err
	}
	log.Debug(fmt.Sprintf("Expanded URL response: %v", expanded))
	return expanded.(string), nil
}
