package storage

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

func init() {
	SupportedStorageTypes["Filesystem"] = new(interface{})
}

type Filesystem struct {
	Root string
	c    uint64
	mu   sync.RWMutex
}

func NewFilesystem(root string) (*Filesystem, error) {
	s := &Filesystem{
		Root: root,
	}
	return s, os.MkdirAll(s.Root, 0744)
}

func (s *Filesystem) Code(url string) string {
	return strconv.FormatUint(s.c, 36)
}

func (s *Filesystem) Save(url string) (string, error) {
	if _, err := validateURL(url); err != nil {
		return "", err
	}

	code := s.Code(url)

	s.mu.Lock()
	err := ioutil.WriteFile(filepath.Join(s.Root, code), []byte(url), 0744)
	if err == nil {
		s.c++
	}
	s.mu.Unlock()

	return code, err
}

// CleanPath removes any path transversal nonsense
func CleanPath(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	return filepath.Clean(path)
}

// Takes a possibly multilevel path and flattens it by dropping any slashes
func FlattenPath(path string, separator string) string {
	return strings.Replace(path, string(os.PathSeparator), separator, -1)
}

func (s *Filesystem) SaveName(rawShort, url string) error {
	short, err := sanitizeShort(rawShort)
	if err != nil {
		return err
	}
	if _, err := validateURL(url); err != nil {
		return err
	}

	short = FlattenPath(CleanPath(short), "_")

	s.mu.Lock()

	if err := ioutil.WriteFile(filepath.Join(s.Root, short), []byte(url), 0744); err == nil {
		s.c++
	}
	s.mu.Unlock()

	return err
}

func (s *Filesystem) Load(rawShort string) (string, error) {
	short, err := sanitizeShort(rawShort)
	if err != nil {
		return "", err
	}

	short = FlattenPath(CleanPath(short), "_")

	s.mu.Lock()
	urlBytes, err := ioutil.ReadFile(filepath.Join(s.Root, short))
	s.mu.Unlock()

	if _, ok := err.(*os.PathError); ok {
		return "", ErrShortNotSet
	}

	return string(urlBytes), err
}
