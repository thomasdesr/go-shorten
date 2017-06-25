package multistorage

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/thomaso-mirodin/go-shorten/storage"
)

type saveTestData struct {
	name           string
	inputShort     string
	inputURL       string
	inputStores    []storage.NamedStorage
	expectedStores []storage.NamedStorage
	expectedErr    error
}

func (stv saveTestData) testFunc(t *testing.T, testedFunc Saver) {
	stores := stv.inputStores

	t.Logf("saving (%q, %q, %q), expecting (%#v)", stv.inputShort, stv.inputURL, stv.inputStores, stv.expectedErr)
	err := testedFunc(context.Background(), stv.inputShort, stv.inputURL, stores)
	t.Logf("got: (%#v)", err)

	if cause := errors.Cause(err); cause != stv.expectedErr {
		t.Errorf("unexpected error: expected(%#v) != actual(%#v)", stv.expectedErr, cause)
	}

	t.Logf("checking equality of stores: %s, %s", stores, stv.expectedStores)
	assert.Equal(t, stv.expectedStores, stores, "stores don't match")
}

func TestSaveOnlyOnceFunc(t *testing.T) {
	storageTestTable := []saveTestData{
		{ // Start with two empty storages and make sure the save only goes into the first one
			name: "Empty to something",
			inputStores: []storage.NamedStorage{
				inmemStorageFromMap(map[string]string{}),
				inmemStorageFromMap(map[string]string{}),
			},
			inputShort: "a",
			inputURL:   "http://A",
			expectedStores: []storage.NamedStorage{
				inmemStorageFromMap(map[string]string{"a": "http://A"}),
				inmemStorageFromMap(map[string]string{}),
			},
			expectedErr: nil,
		},
		{ // We shouldn't modify anything beyond the first successful save
			name: "Doesn't touch the second one",
			inputStores: []storage.NamedStorage{
				inmemStorageFromMap(map[string]string{}),
				inmemStorageFromMap(map[string]string{"B": "http://B"}),
			},
			inputShort: "a",
			inputURL:   "http://A",
			expectedStores: []storage.NamedStorage{
				inmemStorageFromMap(map[string]string{"a": "http://A"}),
				inmemStorageFromMap(map[string]string{"B": "http://B"}),
			},
			expectedErr: nil,
		},
		{ // Doesn't modify any other values than the original
			name: "Prepopulated",
			inputStores: []storage.NamedStorage{
				inmemStorageFromMap(map[string]string{"a": "http://A"}),
				inmemStorageFromMap(map[string]string{}),
			},
			inputShort: "b",
			inputURL:   "http://B",
			expectedStores: []storage.NamedStorage{
				inmemStorageFromMap(map[string]string{
					"a": "http://A",
					"b": "http://B",
				},
				),
				inmemStorageFromMap(map[string]string{}),
			},
			expectedErr: nil,
		},
		{ // Correctly modify a value
			name: "Update in place",
			inputStores: []storage.NamedStorage{
				inmemStorageFromMap(map[string]string{"a": "http://B"}),
				inmemStorageFromMap(map[string]string{}),
			},
			inputShort: "a",
			inputURL:   "http://A",
			expectedStores: []storage.NamedStorage{
				inmemStorageFromMap(map[string]string{"a": "http://A"}),
				inmemStorageFromMap(map[string]string{}),
			},
			expectedErr: nil,
		},
		{ // Test an empty list
			name:           "EmptyList",
			inputStores:    []storage.NamedStorage{},
			inputShort:     "d",
			inputURL:       "http://D",
			expectedStores: []storage.NamedStorage{},
			expectedErr:    ErrEmpty,
		},
	}

	for _, tt := range storageTestTable {
		tt := tt // Scoped copy
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.testFunc(t, saveOnlyOnceFunc)
		})
	}
}

func TestSaveAllFunc(t *testing.T) {
	storageTestTable := []saveTestData{
		{ // Start with two empty storages and make sure they both get modified
			name: "Empty to something",
			inputStores: []storage.NamedStorage{
				inmemStorageFromMap(map[string]string{}),
				inmemStorageFromMap(map[string]string{}),
			},
			inputShort: "a",
			inputURL:   "http://A",
			expectedStores: []storage.NamedStorage{
				inmemStorageFromMap(map[string]string{"a": "http://A"}),
				inmemStorageFromMap(map[string]string{"a": "http://A"}),
			},
			expectedErr: nil,
		},
		{ // Doesn't modify any other values than the original
			name: "Prepopulated",
			inputStores: []storage.NamedStorage{
				inmemStorageFromMap(map[string]string{"B": "http://B"}),
				inmemStorageFromMap(map[string]string{}),
			},
			inputShort: "A",
			inputURL:   "http://A",
			expectedStores: []storage.NamedStorage{
				inmemStorageFromMap(map[string]string{
					"a": "http://A",
					"b": "http://B",
				}),
				inmemStorageFromMap(map[string]string{"a": "http://A"}),
			},
			expectedErr: nil,
		},
		{ // Correctly modify the value in both places
			name: "Update in place",
			inputStores: []storage.NamedStorage{
				inmemStorageFromMap(map[string]string{"a": "http://B"}),
				inmemStorageFromMap(map[string]string{"a": "http://C"}),
			},
			inputShort: "a",
			inputURL:   "http://A",
			expectedStores: []storage.NamedStorage{
				inmemStorageFromMap(map[string]string{"a": "http://A"}),
				inmemStorageFromMap(map[string]string{"a": "http://A"}),
			},
			expectedErr: nil,
		},
		{ // Test an empty list
			name:           "EmptyList",
			inputStores:    []storage.NamedStorage{},
			inputShort:     "d",
			inputURL:       "http://D",
			expectedStores: []storage.NamedStorage{},
			expectedErr:    ErrEmpty,
		},
	}

	for _, tt := range storageTestTable {
		tt := tt // Scoped copy
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tt.testFunc(t, saveAllFunc)
		})
	}
}
