package storage_test

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thomasdesr/go-shorten/storage"
	"github.com/thomasdesr/go-shorten/storage/migrations"
)

func randString(length int) string {
	b := make([]byte, length)
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const maxLetters = len(letters)

	for i := range b {
		b[i] = letters[rand.Intn(maxLetters)]
	}

	return string(b)
}

func saveSomething(s storage.NamedStorage) (short string, long string, err error) {
	short = randString(10)
	long = "http://" + randString(20) + ".com"

	return short, long, s.SaveName(context.Background(), short, long)
}

// type testExternalStorage struct {
// 	globalSetup     func()
// 	perTestSetup    func(testing.TB) storage.Storage
// 	perTestTeardown func(testing.TB) storage.Storage
// 	glboalTeardown  func()
// }

var storageSetups = map[string]func(testing.TB) storage.NamedStorage{
	"Inmem": setupInmemStorage,
	"S3":    setupS3Storage,
	"S3v3Migration": func(t testing.TB) storage.NamedStorage {
		return &migrations.S3v2MigrationStore{setupS3Storage(t).(*storage.S3)}
	},
	"Filesystem": setupFilesystemStorage,
	"Postgres":   setupPostgresStorage,
}

var storageCleanup = map[string]func() error{
	"S3":       cleanupS3Storage,
	"Postgres": cleanupPostgresStorage,
}

func TestMain(m *testing.M) {
	res := m.Run()

	for _, cf := range storageCleanup {
		err := cf()
		if err != nil {
			log.Println("Cleanup error:", err)
		}
	}

	os.Exit(res)
}

func TestNamedStorageSave(t *testing.T) {
	testCode := "test-named-url"
	testURL := "http://google.com"

	for name, setupStorage := range storageSetups {
		setupStorage := setupStorage

		t.Run(name, func(t *testing.T) {
			namedStorage, ok := setupStorage(t).(storage.NamedStorage)

			if assert.True(t, ok, name) {
				err := namedStorage.SaveName(context.Background(), testCode, testURL)
				t.Logf("[%s] namedStorage.SaveName(\"%s\", \"%s\") -> %#v", name, testCode, testURL, err)
				assert.Nil(t, err, name)
			}
		})
	}
}

func TestNamedStorageNormalization(t *testing.T) {
	testCode := "test-named-url"
	testNormalizedCode := "testnamedurl"
	testURL := "http://google.com"

	for name, setupStorage := range storageSetups {
		setupStorage := setupStorage

		t.Run(name, func(t *testing.T) {
			if name == "Postgres" {
				t.Skip("Postgres links are regexes, so shorts are not normalized; near misses get a fuzzy match instead")
			}

			namedStorage, ok := setupStorage(t).(storage.NamedStorage)

			if assert.True(t, ok, name) {
				err := namedStorage.SaveName(context.Background(), testCode, testURL)
				t.Logf("[%s] namedStorage.SaveName(\"%s\", \"%s\") -> %#v", name, testCode, testURL, err)
				assert.Nil(t, err, name)

				a, err := namedStorage.Load(context.Background(), testCode)
				assert.Nil(t, err, name)
				b, err := namedStorage.Load(context.Background(), testNormalizedCode)
				assert.Nil(t, err, name)

				assert.Equal(t, a, b)
			}
		})
	}
}

func TestMissingLoad(t *testing.T) {
	testCode := "non-existent-short-string"

	for name, setupStorage := range storageSetups {
		setupStorage := setupStorage

		t.Run(name, func(t *testing.T) {
			long, err := setupStorage(t).Load(context.Background(), testCode)
			t.Logf("[%s] storage.Load(\"%s\") -> %#v, %#v", name, testCode, long, err)
			assert.NotNil(t, err, name)
			assert.Equal(t, storage.ErrShortNotSet, err, name)
		})
	}
}

func TestLoad(t *testing.T) {
	for name, setupStorage := range storageSetups {
		setupStorage := setupStorage

		t.Run(name, func(t *testing.T) {
			s := setupStorage(t)

			short, long, err := saveSomething(s)
			t.Logf("[%s] saveSomething(s) -> %#v, %#v, %#v", name, short, long, err)
			assert.Nil(t, err, name)

			newLong, err := s.Load(context.Background(), short)
			t.Logf("[%s] storage.Load(\"%s\") -> %#v, %#v", name, short, long, err)
			assert.Nil(t, err, name)

			assert.Equal(t, long, newLong, name)
		})
	}
}

func TestNamedStorageNames(t *testing.T) {
	var shortNames map[string]error = map[string]error{
		"simple":             nil,
		"":                   storage.ErrShortEmpty,
		"1;DROP TABLE names": nil, // A few SQL Injections
		"';DROP TABLE names": nil,
		"œ∑´®†¥¨ˆøπ“‘":       nil, // Fancy Unicode
		"🇺🇸🇦":                nil,
		"社會科學院語學研究所":         nil,
		"ஸ்றீனிவாஸ ராமானுஜன் ஐயங்கார்":         nil,
		"يَّاكَ نَعْبُدُ وَإِيَّاكَ نَسْتَعِي": nil,
		"Po oživlëGromady strojnye tesnâtsâ ":  nil,
		"Powerلُلُصّبُلُلصّبُررً ॣ ॣh ॣ ॣ冗":    nil, // WebOS Crash
	}

	testURL := "http://google.com"

	for storageName, setupStorage := range storageSetups {
		setupStorage := setupStorage

		t.Run(storageName, func(t *testing.T) {
			namedStorage, ok := setupStorage(t).(storage.NamedStorage)
			if !assert.True(t, ok) {
				return
			}

			for short, expectedErr := range shortNames {
				t.Logf("[%s] Saving URL '%s' should result in '%s'", storageName, short, expectedErr)
				err := namedStorage.SaveName(context.Background(), short, testURL)
				assert.Equal(t, expectedErr, err, fmt.Sprintf("[%s] Saving URL '%s' should've resulted in '%s'", storageName, short, expectedErr))

				if err == nil {
					t.Logf("[%s] Loading URL '%s' should result in '%s'", storageName, short, expectedErr)
					url, err := namedStorage.Load(context.Background(), short)
					assert.Equal(t, expectedErr, err, fmt.Sprintf("[%s] Loading URL '%s' should've resulted in '%s'", storageName, short, expectedErr))

					assert.Equal(t, testURL, url, "Saved URL shoud've matched")
				}
			}
		})
	}
}
