package storage_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thomasdesr/go-shorten/storage"
)

func setupInmemStorage(t testing.TB) storage.NamedStorage {
	s, err := storage.NewInmem(8)
	require.Nil(t, err)

	return s
}
