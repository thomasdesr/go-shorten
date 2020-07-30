package multistorage

import (
	"context"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/thomasdesr/go-shorten/storage"
)

// Saver are expected to process a slice of storages and return a result of SaveName(short, url string)
type Saver func(ctx context.Context, short string, url string, stores []storage.NamedStorage) error

func saveAllFunc(ctx context.Context, short string, url string, stores []storage.NamedStorage) error {
	if len(stores) == 0 {
		return ErrEmpty
	}

	errs := new(multierror.Error)
	for _, store := range stores {
		err := store.SaveName(ctx, short, url)

		if err != nil {
			multierror.Append(
				errs,
				errors.Wrapf(err, "failed to save %q to %q", short, store),
			)
		}
	}

	return errs.ErrorOrNil()
}

// saveOnlyOnceFunc
func saveOnlyOnceFunc(ctx context.Context, short string, url string, stores []storage.NamedStorage) error {
	if len(stores) == 0 {
		return ErrEmpty
	}

	errs := new(multierror.Error)
	for _, store := range stores {
		err := store.SaveName(ctx, short, url)

		if err == nil {
			return nil
		}

		multierror.Append(errs, err)
	}

	return errors.Wrap(
		errs.ErrorOrNil(),
		"failed to save to only one store",
	)
}
