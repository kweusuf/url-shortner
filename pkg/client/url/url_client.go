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
}

type urlClient struct {
}

func MakeURLClient() URLClient {
	return &urlClient{}
}

// ShortenURL implements URLClient.
func (client *urlClient) ShortenURL(ctx context.Context, url string) (interface{}, error) {
	log.Debug(fmt.Sprintf("In ShortenURL method with URL: %s", url))

	shortCode, err := urlUtil.ShortenURL(ctx, url)
	if err != nil {
		return nil, err
	}

	shortenedURL := fmt.Sprintf("https://my-domain.com/%s", shortCode)
	return shortenedURL, nil
}
