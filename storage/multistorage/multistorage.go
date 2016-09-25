package multistorage

import (
	"github.com/pkg/errors"
	"github.com/thomaso-mirodin/go-shorten/storage"
)

func init() {
	storage.SupportedStorageTypes["MultiStorage"] = new(interface{})
}

// MultiStorage is a storage.NamedStorage that will allow you to interact with multiple underlying storage.NamedStorages.
type MultiStorage struct {
	stores []storage.NamedStorage
	loader Loader
}

func New(stores []storage.NamedStorage, opts ...MultiStorageOption) (*MultiStorage, error) {
	m := &MultiStorage{
		stores: stores,
		loader: loadFirstFunc,
	}

	for _, opt := range opts {
		if err := opt(m); err != nil {
			return nil, err
		}
	}

	if err := m.validateStore(); err != nil {
		return nil, errors.Wrap(err, "store failed to validate")
	}

	return m, nil
}

func Simple(stores ...storage.NamedStorage) (*MultiStorage, error) {
	return New(stores, LoadFirst())
}

var ErrEmpty = errors.New("MultiStorage has no underlying stores")

func (s *MultiStorage) validateStore() error {
	if len(s.stores) == 0 {
		return ErrEmpty
	}

	return nil
}

// Load with a basic MultiStorage will query the underlying storages (in order) returning when either a response or error is encountered, only returning an ErrShortNotSet when all underlying storages have been exhausted.
func (s *MultiStorage) Load(short string) (string, error) {
	if err := s.validateStore(); err != nil {
		return "", errors.Wrap(err, "failed to validate underlying store")
	}

	return s.loader(short, s.stores)
}

// // SaveName will return the first successful insure that all
// func (s *MultiStorage) SaveName(short string, long string) error {
// 	for _, store := range *s {
// 		err := store.SaveName(short, long)
// 		if err != nil {
// 			return errors.Wrapf(err, "failed to save %q to %q", short, store)
// 		}
// 	}
// }
