package storage_test

import (
	"context"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thomasdesr/go-shorten/internal/errorutil"
	"github.com/thomasdesr/go-shorten/storage"
)

// connectString points at a Postgres with the flyway/sql migrations applied, e.g.
// "postgres://postgres@localhost/?sslmode=disable". Postgres tests skip when it is unset.
var connectString = os.Getenv("POSTGRES_TEST_URL")

func TestPostgresLoadCountsHitsForLoadedShort(t *testing.T) {
	ctx := context.Background()
	p := setupPostgresStorage(t).(*storage.Postgres)

	// Two shorts sharing one URL give two links rows joined to a single urls row,
	// so a hit recorded against any id other than each short's own link cannot
	// produce these counts.
	url := "http://" + randString(20) + ".com"
	once, twice := randString(10), randString(10)
	require.NoError(t, p.SaveName(ctx, once, url))
	require.NoError(t, p.SaveName(ctx, twice, url))

	for _, short := range []string{once, twice, twice} {
		_, err := p.Load(ctx, short)
		require.NoError(t, err)
	}

	top, err := p.TopNForPeriod(ctx, 1000, 1)
	require.NoError(t, err)

	hits := map[string]int{}
	for _, r := range top {
		hits[r.Link] = r.HitCount
	}
	assert.Equal(t, 1, hits[once], "hits for %q", once)
	assert.Equal(t, 2, hits[twice], "hits for %q", twice)
}

func setupPostgresStorage(t testing.TB) storage.NamedStorage {
	if connectString == "" {
		t.Skip("POSTGRES_TEST_URL is unset")
	}

	p, err := storage.NewPostgres(connectString)
	if err != nil {
		t.Fatal(errorutil.Wrap(err, "failed to create storage.Postgres"))
	}

	return p
}

func cleanupPostgresStorage() error {
	if connectString == "" {
		return nil
	}

	dbx, err := sqlx.Connect("postgres", connectString)
	if err != nil {
		return errorutil.Wrap(err, "failed to connect to Postgres")
	}

	_, err = dbx.Exec("DELETE FROM LINKS_USAGE; DELETE FROM LINKS; DELETE FROM URLS;")
	return err
}
