package storage_test

import (
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/thomasdesr/go-shorten/storage"
)

var connectString = "postgres://postgres@localhost/?sslmode=disable"

func setupPostgresStorage(t testing.TB) storage.NamedStorage {
	p, err := storage.NewPostgres(connectString)
	if err != nil {
		t.Fatal(errors.Wrap(err, "failed to create storage.Postgres"))
	}

	return p
}

func cleanupPostgresStorage() error {
	dbx, err := sqlx.Connect("postgres", connectString)
	if err != nil {
		return errors.Wrap(err, "failed to connect to Postgres")
	}

	_, err = dbx.Exec("DELETE FROM LINKS; DELETE FROM URLS;")
	return err
}
