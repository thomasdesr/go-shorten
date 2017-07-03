package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

type Inmem struct {
	RandLength int

	m  map[string]string
	mu sync.RWMutex
}

func (s Inmem) String() string {
	j := struct {
		RandLength int
		InnerMap   map[string]string
	}{s.RandLength, s.m}

	b, err := json.Marshal(j)
	if err != nil {
		return fmt.Sprintf("%#v", s)
	}

	return string(b)
}

func NewInmem(randLength int) (*Inmem, error) {
	s := &Inmem{
		RandLength: randLength,

		m: make(map[string]string),
	}
	return s, nil
}

func NewInmemFromMap(randLength int, initialShorts map[string]string) (*Inmem, error) {
	s, _ := NewInmem(randLength)

	for k, v := range initialShorts {
		if err := s.SaveName(context.Background(), k, v); err != nil {
			return nil, errors.Wrap(err, "failed to save initial short")
		}
	}

	return s, nil
}

func (s *Inmem) SaveName(ctx context.Context, rawShort string, url string) error {
	short, err := sanitizeShort(rawShort)
	if err != nil {
		return err
	}
	if _, err := validateURL(url); err != nil {
		return err
	}

	s.mu.Lock()
	s.m[short] = url
	s.mu.Unlock()
	return nil
}

func (s *Inmem) Load(ctx context.Context, rawShort string) (string, error) {
	short, err := sanitizeShort(rawShort)
	if err != nil {
		return "", err
	}

	s.mu.Lock()
	url, ok := s.m[short]
	s.mu.Unlock()
	if !ok {
		return "", ErrShortNotSet
	}

	return url, nil
}
