package storage

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"path"
	"time"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
)

func init() {
	SupportedStorageTypes["S3"] = new(interface{})
}

type S3 struct {
	Bucket         *s3.Bucket
	hashFunc       func(string) string
	storageVersion string
}

func NewS3(auth aws.Auth, region aws.Region, bucketName string) (*S3, error) {
	s := &S3{
		Bucket: s3.New(auth, region).Bucket(bucketName),
		hashFunc: func(s string) string {
			h := sha256.Sum256([]byte(s))
			return hex.EncodeToString(h[:])
		},
		storageVersion: "v2",
	}

	_, err := s.Bucket.List("/", "", "", 1)
	if s3err, ok := err.(*s3.Error); ok && s3err.Code == "NoSuchBucket" {
		err = s.Bucket.PutBucket(s3.BucketOwnerFull)
	}

	return s, err
}

func (s *S3) saveKey(short, url string) (err error) {
	hashedShort := s.hashFunc(short)

	err = s.Bucket.Put(
		path.Join(s.storageVersion, hashedShort, "long"),
		[]byte(url),
		"text/plain",
		s3.BucketOwnerFull,
		s3.Options{},
	)
	if err != nil {
		return err
	}

	err = s.Bucket.Put(
		path.Join(s.storageVersion, hashedShort, "short"),
		[]byte(short),
		"text/plain",
		s3.BucketOwnerFull,
		s3.Options{},
	)
	if err != nil {
		return err
	}

	changeLog, err := json.Marshal(
		struct {
			URL  string
			User string
		}{
			url,
			"TODO",
		},
	)
	if err != nil {
		return fmt.Errorf("unable to format change history: %v", err)
	}

	err = s.Bucket.Put(
		path.Join(s.storageVersion, hashedShort, "change_history", time.Now().Format(time.RFC3339Nano)),
		changeLog,
		"application/json",
		s3.BucketOwnerFull,
		s3.Options{},
	)
	if err != nil {
		return fmt.Errorf("unable to save change history: %v", err)
	}

	return nil
}

func (s *S3) Save(url string) (string, error) {
	if _, err := validateURL(url); err != nil {
		return "", err
	}

	for i := 0; i < 10; i++ {
		short := getRandomString(8)
		pathToShort := path.Join(s.storageVersion, s.hashFunc(short))

		if exists, err := s.Bucket.Exists(pathToShort); !exists && err == nil {
			return short, s.saveKey(short, url)
		}
	}

	return "", ErrShortExhaustion
}

func (s *S3) SaveName(rawShort string, url string) error {
	short, err := sanitizeShort(rawShort)
	if err != nil {
		return err
	}
	if _, err := validateURL(url); err != nil {
		return err
	}

	return s.saveKey(short, url)
}

func (s *S3) Load(rawShort string) (string, error) {
	short, err := sanitizeShort(rawShort)
	if err != nil {
		return "", err
	}

	url, err := s.Bucket.Get(path.Join(s.storageVersion, s.hashFunc(short), "long"))
	if s3err, ok := err.(*s3.Error); ok && s3err.Code == "NoSuchKey" {
		return "", ErrShortNotSet
	}
	return string(url), err
}
