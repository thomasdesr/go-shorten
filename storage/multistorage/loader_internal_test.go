package multistorage

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/thomaso-mirodin/go-shorten/storage"
)

func inmemStorageFromMap(inputs map[string]string) *storage.Inmem {
	s, err := storage.NewInmemFromMap(8, inputs)
	if err != nil {
		panic(err)
	}

	return s
}

func TestLoadFirstFunc(t *testing.T) {
	inputs := []map[string]string{
		{"a": "http://A"},
		{"b": "http://B"},
		{"a": "http://C"}, // Duplicate key to test ordering
	}

	storageTestTable := []struct {
		name         string
		stores       []storage.NamedStorage
		inputShort   string
		expectedLong string
		expectedErr  error
	}{
		{ // Load the inputs in the above order and try to get the first "a" value back ("http://A")
			name: "Simple",
			stores: []storage.NamedStorage{
				inmemStorageFromMap(inputs[0]),
				inmemStorageFromMap(inputs[1]),
				inmemStorageFromMap(inputs[2]),
			},
			inputShort:   "a",
			expectedLong: "http://A",
			expectedErr:  nil,
		},
		{ // Test the the inputs in the opposite order so we should get the later "a" value ("http://C")
			name: "ReverseSimple",
			stores: []storage.NamedStorage{
				inmemStorageFromMap(inputs[2]),
				inmemStorageFromMap(inputs[1]),
				inmemStorageFromMap(inputs[0]),
			},
			inputShort:   "a",
			expectedLong: "http://C",
			expectedErr:  nil,
		},
		{ // Test for a missing value
			name: "MissingValue",
			stores: []storage.NamedStorage{
				inmemStorageFromMap(inputs[0]),
				inmemStorageFromMap(inputs[1]),
				inmemStorageFromMap(inputs[2]),
			},
			inputShort:   "c",
			expectedLong: "",
			expectedErr:  storage.ErrShortNotSet,
		},
		{ // Test an empty list
			name:         "EmptyList",
			stores:       []storage.NamedStorage{},
			inputShort:   "d",
			expectedLong: "",
			expectedErr:  ErrEmpty,
		},
	}

	for _, tt := range storageTestTable {
		tt := tt // Scoped copy
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			t.Logf("querying for %q, expecting (%q,%#v)", tt.inputShort, tt.expectedLong, tt.expectedErr)
			long, err := loadFirstFunc(tt.inputShort, tt.stores)
			t.Logf("got: (%q, %#v)", long, err)
			if cause := errors.Cause(err); cause != tt.expectedErr {
				t.Errorf("unexpected error: expected(%#v) != actual(%#v)", tt.expectedErr, cause)
			}

			if long != tt.expectedLong {
				t.Errorf("unexpected long: expected(%q) != actual(%q)", tt.expectedLong, long)
			}
		})
	}
}

func TestLoadCompareAllResultsFunc(t *testing.T) {
	inputs := []map[string]string{
		{"a": "http://A"},
		{"a": "http://A"},
		{"b": "http://B"},
		{"b": "http://C"}, // Duplicate key to test equality checks
	}

	storageTestTable := []struct {
		name         string
		stores       []storage.NamedStorage
		inputShort   string
		expectedLong string
		expectedErr  error
	}{
		{ // Test a load with only one value so should return the correct value without error
			name: "SingleKeySingleValue",
			stores: []storage.NamedStorage{
				inmemStorageFromMap(inputs[0]),
				inmemStorageFromMap(inputs[1]),
			},
			inputShort:   "a",
			expectedLong: "http://A",
			expectedErr:  nil,
		},
		{ // Test for a value with multiple different answers for the same key, this should return no value with an error
			name: "SingleKeyMultipleValues",
			stores: []storage.NamedStorage{
				inmemStorageFromMap(inputs[2]),
				inmemStorageFromMap(inputs[3]),
			},
			inputShort:   "b",
			expectedLong: "",
			expectedErr:  ErrUnexpectedMultipleAnswers,
		},
		{ // Test a load with multiple keys and multiple values, this should fail with an error
			name: "MultipleKeysMultipleValues",
			stores: []storage.NamedStorage{
				inmemStorageFromMap(inputs[0]),
				inmemStorageFromMap(inputs[1]),
				inmemStorageFromMap(inputs[2]),
				inmemStorageFromMap(inputs[3]),
			},
			inputShort:   "a",
			expectedLong: "",
			expectedErr:  ErrUnexpectedMultipleAnswers,
		},
		{ // Test for a missing value
			name: "MissingValue",
			stores: []storage.NamedStorage{
				inmemStorageFromMap(inputs[0]),
				inmemStorageFromMap(inputs[1]),
				inmemStorageFromMap(inputs[2]),
				inmemStorageFromMap(inputs[3]),
			},
			inputShort:   "c",
			expectedLong: "",
			expectedErr:  storage.ErrShortNotSet,
		},
		{ // Test an empty list
			name:         "EmptyList",
			stores:       []storage.NamedStorage{},
			inputShort:   "d",
			expectedLong: "",
			expectedErr:  ErrEmpty,
		},
	}

	for _, tt := range storageTestTable {
		tt := tt // Scoped copy
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			t.Logf("querying for %q, expecting (%q,%#v)", tt.inputShort, tt.expectedLong, tt.expectedErr)
			long, err := loadCompareAllResultsFunc(tt.inputShort, tt.stores)
			t.Logf("got: (%q, %#v)", long, err)
			if cause := errors.Cause(err); cause != tt.expectedErr {
				t.Errorf("unexpected error: expected(%#v) != actual(%#v)", tt.expectedErr, cause)
			}

			if long != tt.expectedLong {
				t.Errorf("unexpected long: expected(%q) != actual(%q)", tt.expectedLong, long)
			}
		})
	}
}
