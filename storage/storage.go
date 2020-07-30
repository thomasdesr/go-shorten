// Package storages allows multiple implementation on how to store URLs as shorter names and retrieve them later.
package storage

import (
	"context"
	"errors"
	"net/url"
	"strings"
)

type Storage interface {
	// Load(ctx, string) takes a short URL and returns the original full URL by retrieving it from storage
	Load(ctx context.Context, short string) (string, error)
}

type NamedStorage interface {
	Storage
	// SaveName takes a short and a url and saves the name to use for saving a url
	SaveName(ctx context.Context, short string, url string) error
}

type SearchableStorage interface {
	Storage
	// Search takes a search term and returns a number of possible shorts
	Search(ctx context.Context, searchTerm string) ([]SearchResult, error)
}

type TopN interface {
	Storage
	// TopNForPeriod returns the most visited shorts in the last N days
	TopNForPeriod(ctx context.Context, n int, days int) ([]TopNResult, error)
}

var (
	ErrURLEmpty   = errors.New("provided URL is of zero length")
	ErrShortEmpty = errors.New("provided short name is of zero length")

	ErrURLNotAbsolute = errors.New("provided URL is not an absolute URL")

	ErrShortNotSet = errors.New("storage layer doens't have a URL for that short code")

	ErrFuzzyMatchFound = errors.New("fuzzy match found")

	ErrNoResults = errors.New("No search results found")
)

func validateShort(short string) error {
	if short == "" {
		return ErrShortEmpty
	}

	return nil
}

func validateURL(rawURL string) (*url.URL, error) {
	if rawURL == "" {
		return nil, ErrShortEmpty
	}

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}

	if !parsedURL.IsAbs() {
		return nil, ErrURLNotAbsolute
	}

	return parsedURL, nil
}

var normalizingReplacer = strings.NewReplacer(
	" ", "",
	"-", "",
	"_", "",
	"/", "",
	",", "",
	".", "",
)

func sanitizeShort(rawShort string) (string, error) {
	short := normalizingReplacer.Replace(strings.ToLower(rawShort))

	return short, validateShort(short)
}
