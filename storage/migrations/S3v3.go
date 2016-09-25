package migrations

import "github.com/thomaso-mirodin/go-shorten/storage"

// S3v2MigrationStore helps the migration from the v2 version of the store to
// the v3 version of the store that also stores the original short code
type S3v2MigrationStore struct {
	*storage.S3
}

func (s *S3v2MigrationStore) Load(short string) (long string, err error) {
	long, err = s.S3.Load(short)
	if err != nil {
		return
	}

	err = s.S3.SaveName(short, long)
	return
}
