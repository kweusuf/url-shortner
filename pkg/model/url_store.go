package model

import "time"

type URLStoreEntry struct {
	CreatedTimestamp time.Time `json:"created_timestamp"`
	OriginalURL      string    `json:"original_url"`
	ShortenedURL     string    `json:"shortened_url"`
	ShortCode        string    `json:"short_code"`
	ExpirationDate   time.Time `json:"expiration_date"`
}
