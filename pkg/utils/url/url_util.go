package url

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/kweusuf/url-shortner/pkg/constants"
	"github.com/kweusuf/url-shortner/pkg/model"
	"github.com/kweusuf/url-shortner/pkg/utils/log"
)

var (
	urlStore = make(map[string]model.URLStoreEntry) // Store for original URLs and their shortened versions
	mu       sync.Mutex
	counter  int
)

func ShortenURL(ctx context.Context, url string) (interface{}, error) {
	log.Debug(fmt.Sprintf("In ShortenURL Util method with URL: %s", url))
	mu.Lock()
	defer mu.Unlock()

	// Check if already shortened
	val, ok := urlStore[url]
	// If the key exists
	if ok {
		log.Debug(fmt.Sprintf("URL already shortened: %s -> %s", url, val.ShortenedURL))
		return val.ShortenedURL, nil
	}
	// Generate a simple short code
	counter++
	shortCode := fmt.Sprintf("short%d", counter)
	urlStore[url] = model.URLStoreEntry{
		CreatedTimestamp: time.Now(),
		OriginalURL:      url,
		ShortenedURL:     shortCode,
		ShortCode:        shortCode,
		ExpirationDate:   time.Now().Add(constants.ExpirationInterval),
	}

	return shortCode, nil
}

func ExpandURL(ctx context.Context, url string) (interface{}, error) {
	log.Debug(fmt.Sprintf("In ExpandURL Util method with URL: %s", url))

	// Check if the URL exists in the store
	for short, orig := range urlStore {
		if short == url {
			return orig, nil
		}
	}
	return nil, fmt.Errorf("URL not found: %s", url)
}

func CleanUpStaleEntries() {
	for key, elem := range urlStore {
		if elem.ExpirationDate.Before(time.Now()) {
			delete(urlStore, key)
		}
	}
}

func ExecuteCleanUp() {
	ticker := time.NewTicker(constants.CleanupInterval)
	defer ticker.Stop()

	for {
		<-ticker.C // Wait for the next tick
		CleanUpStaleEntries()
		log.Info(fmt.Sprintf("Cleanup task executed at: %s", time.Now()))
	}
}
