package multistorage

import (
	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/thomaso-mirodin/go-shorten/storage"
)

// Saver are expected to process a slice of storages and return a result of SaveName(short, url string)
type Saver func(short string, url string, stores []storage.NamedStorage) error

func saveAllFunc(short string, url string, stores []storage.NamedStorage) error {
	if len(stores) == 0 {
		return ErrEmpty
	}

	var errs *multierror.Error
	for _, store := range stores {
		err := store.SaveName(short, url)

		if err != nil {
			multierror.Append(
				errs,
				errors.Wrapf(err, "failed to save %q to %q", short, store),
			)
		}
	}

	return errs.ErrorOrNil()
}

func saveBestEffortFunc(short string, url string, stores []storage.NamedStorage) error {
	if len(stores) == 0 {
		return ErrEmpty
	}

	for _, store := range stores {
		_ = store.SaveName(short, url)

	}

	return nil
}
