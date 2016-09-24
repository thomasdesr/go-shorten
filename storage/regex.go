package storage

import (
	"errors"
	"regexp"
	"strings"
	"sync"
)

func init() {
	SupportedStorageTypes["Regex"] = new(interface{})
}

type remap struct {
	Regex       *regexp.Regexp
	Replacement string
}

type Regex struct {
	remaps []remap
	mu     sync.RWMutex
}

func NewRegexFromList(redirects map[string]string) (*Regex, error) {
	remaps := make([]remap, 0, len(redirects))

	for regexString, redirect := range redirects {
		r, err := regexp.Compile(regexString)
		if err != nil {
			return nil, err
		}
		remaps = append(remaps, remap{
			Regex:       r,
			Replacement: redirect,
		})
	}

	return &Regex{
		remaps: remaps,
	}, nil
}

func (r *Regex) Load(short string) (string, error) {
	// Regex intentionally doesn't do sanitization, each regex can have whatever flexability it wants
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, remap := range r.remaps {
		if remap.Regex.MatchString(short) {
			return remap.Regex.ReplaceAllString(short, remap.Replacement), nil
		}
	}

	return "", ErrShortNotSet
}

var ErrRegexIncorrectFormat = errors.New("regex format doens't match the format '/.+/'")

func (r *Regex) SaveName(short string, long string) (string, error) {
	// Validate that the short starts with and ends with a '/'
	if !strings.HasPrefix(short, "/") {
		return "", errors.Wrapf(ErrRegexIncorrectFormat, "%q doesn't start with a slash", regexp.QuoteMeta(short))
	}
	if !strings.HasSuffix(short, "/") {
		return "", errors.Wrapf(ErrRegexIncorrectFormat, "%q doesn't end with a slash", regexp.QuoteMeta(short))
	}

	r.mu.Lock()
	defer r.mu.Unlock()
	r.remaps = append() // This is hard, how do I make sure there aren't collisions?
}
