package storage

import (
	"fmt"
	"regexp"
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

func (r Regex) Load(short string) (string, error) {
	// Regex intentionally doesn't do sanitization, each regex can have whatever flexability it wants

	for _, remap := range r.remaps {
		if remap.Regex.MatchString(short) {
			return remap.Regex.ReplaceAllString(short, remap.Replacement), nil
		}
	}

	return "", ErrShortNotSet
}

func (r Regex) SaveName(short string, long string) (string, error) {
	// Regex intentionally doesn't do sanitization, each regex can have whatever flexability it wants

	return "", fmt.Errorf("regex doesn't yet support saving after creation")
}
