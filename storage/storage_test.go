package storage_test

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thomaso-mirodin/go-shorten/storage"
	"github.com/thomaso-mirodin/go-shorten/storage/migrations"
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

func saveSomething(s storage.Storage) (short string, long string, err error) {
	unnamed, unnamedOk := s.(storage.UnnamedStorage)
	named, namedOk := s.(storage.NamedStorage)

	short = randString(10)
	long = "http://" + randString(20) + ".com"

	if namedOk {
		err := named.SaveName(short, long)
		return short, long, err
	} else if unnamedOk {
		short, err := unnamed.Save(long)
		return short, long, err
	} else {
		return "", "", fmt.Errorf("Storage isn't named or unnamed, can't save anything")
	}
}

// type testExternalStorage struct {
// 	globalSetup     func()
// 	perTestSetup    func(testing.TB) storage.Storage
// 	perTestTeardown func(testing.TB) storage.Storage
// 	glboalTeardown  func()
// }

var storageSetups = map[string]func(testing.TB) storage.Storage{
	"Inmem": setupInmemStorage,
	"S3":    setupS3Storage,
	"S3v3Migration": func(t testing.TB) storage.Storage {
		return &migrations.S3v2MigrationStore{setupS3Storage(t).(*storage.S3)}
	},
	"Filesystem": setupFilesystemStorage,
}

var storageCleanup = map[string]func() error{
	"S3": cleanupS3Storage,
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

func TestUnnamedStorageSave(t *testing.T) {
	testURL := "http://google.com"

	for name, setupStorage := range storageSetups {
		setupStorage := setupStorage

		t.Run(name, func(t *testing.T) {
			unnamedStorage, ok := setupStorage(t).(storage.UnnamedStorage)

			if assert.True(t, ok, name) {
				code, err := unnamedStorage.Save(testURL)
				t.Logf("[%s] unnamedStorage.Save(\"%s\") -> %#v", name, testURL, code)
				assert.Nil(t, err, name)
			}
		})
	}
}

func TestNamedStorageSave(t *testing.T) {
	testCode := "test-named-url"
	testURL := "http://google.com"

	for name, setupStorage := range storageSetups {
		setupStorage := setupStorage

		t.Run(name, func(t *testing.T) {
			namedStorage, ok := setupStorage(t).(storage.NamedStorage)

			if assert.True(t, ok, name) {
				err := namedStorage.SaveName(testCode, testURL)
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
			namedStorage, ok := setupStorage(t).(storage.NamedStorage)

			if assert.True(t, ok, name) {
				err := namedStorage.SaveName(testCode, testURL)
				t.Logf("[%s] namedStorage.SaveName(\"%s\", \"%s\") -> %#v", name, testCode, testURL, err)
				assert.Nil(t, err, name)

				a, err := namedStorage.Load(testCode)
				assert.Nil(t, err, name)
				b, err := namedStorage.Load(testNormalizedCode)
				assert.Nil(t, err, name)

				assert.Equal(t, a, b)
			}
		})
	}
}

func TestMissingLoad(t *testing.T) {
	testCode := "non-existant-short-string"

	for name, setupStorage := range storageSetups {
		setupStorage := setupStorage

		t.Run(name, func(t *testing.T) {
			long, err := setupStorage(t).Load(testCode)
			t.Logf("[%s] storage.Load(\"%s\") -> %#v, %#v", name, testCode, long, err)
			assert.NotNil(t, err, name)
			assert.Equal(t, err, storage.ErrShortNotSet, name)
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

			newLong, err := s.Load(short)
			t.Logf("[%s] storage.Load(\"%s\") -> %#v, %#v", name, short, long, err)
			assert.Nil(t, err, name)

			assert.Equal(t, long, newLong, name)
		})
	}
}

func TestNamedStorageNames(t *testing.T) {
	var shortNames map[string]error = map[string]error{
		"simple":                               nil,
		"":                                     storage.ErrShortEmpty,
		"1;DROP TABLE names":                   nil, // A few SQL Injections
		"';DROP TABLE names":                   nil,
		"œ∑´®†¥¨ˆøπ“‘":                         nil, // Fancy Unicode
		"🇺🇸🇦":                                  nil,
		"社會科學院語學研究所":                           nil,
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

			for short, e := range shortNames {
				t.Logf("[%s] Saving URL '%s' should result in '%s'", storageName, short, e)
				err := namedStorage.SaveName(short, testURL)
				assert.Equal(t, err, e, fmt.Sprintf("[%s] Saving URL '%s' should've resulted in '%s'", storageName, short, e))

				if err == nil {
					t.Logf("[%s] Loading URL '%s' should result in '%s'", storageName, short, e)
					url, err := namedStorage.Load(short)
					assert.Equal(t, err, e, fmt.Sprintf("[%s] Loading URL '%s' should've resulted in '%s'", storageName, short, e))

					assert.Equal(t, url, testURL, "Saved URL shoud've matched")
				}
			}
		})
	}
}
