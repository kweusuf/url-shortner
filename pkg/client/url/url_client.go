package url

import (
	"context"
	"fmt"
	"sync"

	"github.com/kweusuf/url-shortner/pkg/utils/log"
	urlUtil "github.com/kweusuf/url-shortner/pkg/utils/url"
)

var (
	urlStore = make(map[string]string)
	mu       sync.Mutex
	counter  int
)

type URLClient interface {
	ShortenURL(ctx context.Context, url string) (interface{}, error)
	ExpandURL(ctx context.Context, url string) (interface{}, error)
}

type urlClient struct {
}

func MakeURLClient() URLClient {
	return &urlClient{}
}

// ShortenURL implements URLClient.
func (client *urlClient) ShortenURL(ctx context.Context, url string) (interface{}, error) {
	log.Debug(fmt.Sprintf("In ShortenURL method with URL: %s", url))

	shortenedURL, err := urlUtil.ShortenURL(ctx, url)
	if err != nil {
		return nil, err
	}

	return shortenedURL, nil
}

// ExpandURL implements URLClient.
func (client *urlClient) ExpandURL(ctx context.Context, url string) (interface{}, error) {
	log.Debug(fmt.Sprintf("In ExpandURL method with URL: %s", url))

	originalURL, err := urlUtil.ExpandURL(ctx, url)
	if err != nil {
		return nil, err
	}

	if originalURL == nil {
		return nil, fmt.Errorf("URL not found: %s", url)
	}

	return originalURL, nil
}
