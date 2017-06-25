package multistorage_test

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/thomaso-mirodin/go-shorten/storage"
	"github.com/thomaso-mirodin/go-shorten/storage/multistorage"
)

func inmemStorageFromMap(inputs map[string]string) *storage.Inmem {
	s, err := storage.NewInmemFromMap(8, inputs)
	if err != nil {
		panic(err)
	}

	return s
}

func TestNoBackendsNew(t *testing.T) {
	expectedErr := multistorage.ErrEmpty

	_, err := multistorage.New(
		[]storage.NamedStorage{},
	)

	if cause := errors.Cause(err); expectedErr != cause {
		t.Fatal("making the multistorage without any underlying storages should've failed")
	}
}

func TestSingleBackend(t *testing.T) {
	inputShort := "abc"
	inputLong := "http://def"

	m, err := multistorage.Simple(
		inmemStorageFromMap(map[string]string{}),
	)
	if err != nil {
		t.Fatalf("failed to create multistorage because %q", err)
	}

	t.Logf("Saving %q->%q", inputShort, inputLong)
	if err := m.SaveName(context.Background(), inputShort, inputLong); err != nil {
		t.Fatalf("error saving %q->%q into the store", inputShort, inputLong)
	}
	t.Logf("Got: %v", err)

	t.Logf("Loading %q", inputShort)
	long, err := m.Load(context.Background(), inputShort)
	t.Logf("Got: %q, %v", long, err)
	if err != nil {
		t.Fatalf("error loading value that should exist: %q", err)
	}

	if long != inputLong {
		t.Fatalf("returned incorrect value from the underlying stores: %q != %q", long, inputLong)
	}
}

func TestMultipleBackendLoad(t *testing.T) {
	inputShorts := []map[string]string{
		{"a": "http://A"},
		{"b": "http://B"},
		{"c": "http://C"},
	}

	m, err := multistorage.Simple(
		inmemStorageFromMap(inputShorts[0]),
		inmemStorageFromMap(inputShorts[1]),
		inmemStorageFromMap(inputShorts[2]),
	)
	if err != nil {
		t.Fatal("failed creating multistorage", err)
	}

	for _, input := range inputShorts {
		for inputShort, expectedLong := range input {
			t.Logf("Loading %q", inputShort)
			long, err := m.Load(context.Background(), inputShort)
			t.Logf("Got: %q, %v", long, err)
			if err != nil {
				t.Errorf("loading %q returned an error: %q", inputShort, err)
			}

			if long != expectedLong {
				t.Errorf("%q != %q", long, expectedLong)
			}
		}
	}
}

// func TestQuickSingleBackend(t *testing.T) {
// 	f := func(shortens map[string]string) bool {
// 		m, err := multistorage.New(
// 			[]storage.NamedStorage{
// 				inmemStorageFromMap(shortens),
// 			},
// 		)
// 		if err != nil {
// 			t.Fatal(errors.Wrap(err, "failed to create multistorage"))
// 		}

// 		for k, v := range shortens {
// 			s, err := m.Load(k)
// 			if err != nil {
// 				return false
// 			}

// 			if s != v {
// 				return false
// 			}
// 		}
// 		return true
// 	}

// 	if err := quick.Check(f, nil); err != nil {
// 		t.Error(err)
// 	}
// }
