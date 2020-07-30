package multistorage

import (
	"context"

	"github.com/pkg/errors"
	"github.com/thomasdesr/go-shorten/storage"
)

// Loaders are expected to process the slice of stores and return the result of Load(short) from one of them. Should return ErrEmpty if stores is empty
type Loader func(ctx context.Context, short string, stores []storage.NamedStorage) (string, error)

func loadFirstFunc(ctx context.Context, short string, stores []storage.NamedStorage) (string, error) {
	if len(stores) == 0 {
		return "", ErrEmpty
	}

	for _, store := range stores {
		long, err := store.Load(ctx, short)
		if err == storage.ErrShortNotSet {
			continue
		}

		return long, err
	}
	return "", storage.ErrShortNotSet
}

var ErrUnexpectedMultipleAnswers = errors.New("MultiStorage: results returned were not the same")

func loadCompareAllResultsFunc(ctx context.Context, short string, stores []storage.NamedStorage) (string, error) {
	if len(stores) == 0 {
		return "", ErrEmpty
	}

	results := make([]loadResult, 0, len(stores))
	for _, store := range stores {
		s, err := store.Load(ctx, short)

		results = append(results, loadResult{s, err})
	}

	if !allSameLoadResults(results) {
		return "", errors.Wrapf(ErrUnexpectedMultipleAnswers, "%#v", results)
	}

	res := results[0]
	if res.err == storage.ErrShortNotSet {
		return "", storage.ErrShortNotSet
	}

	if res.long == "" && res.err == nil {
		panic("something went very wrong, all of the backends returned empty strings for longs and no error")
	}

	return res.long, res.err
}

type loadResult struct {
	long string
	err  error
}

func allSameLoadResults(res []loadResult) bool {
	for i := 1; i < len(res); i++ {
		if res[i] != res[0] {
			return false
		}
	}
	return true
}
