package storage_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/thomaso-mirodin/go-shorten/storage"
)

var testBucket string = "go-shortener-test"

func setupS3Storage(t testing.TB) storage.NamedStorage {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")})
	if err != nil {
		t.Fatal(errors.Wrap(err, "failed to create aws session"))
	}

	s, err := storage.NewS3(sess, testBucket)
	if err != nil {
		t.Fatal(errors.Wrap(err, "failed to create storage.S3"))
	}

	return s
}

func cleanupS3Storage() error {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-west-2")})
	if err != nil {
		return errors.Wrap(err, "failed to create aws session")
	}

	s, err := storage.NewS3(sess, testBucket)
	if err != nil {
		return errors.Wrap(err, "failed to create storage.S3")
	}

	// List all the objects in the test bucket and delete them
	if err := s.Client.ListObjectsV2Pages(
		&s3.ListObjectsV2Input{
			Bucket: aws.String(s.BucketName),
		},
		func(page *s3.ListObjectsV2Output, lastPage bool) (shouldContinue bool) {
			for _, obj := range page.Contents {
				s.Client.DeleteObject(&s3.DeleteObjectInput{
					Bucket: aws.String(s.BucketName),
					Key:    obj.Key,
				})
			}

			return true
		},
	); err != nil {
		return errors.Wrap(err, "failed to list and delete objects in bucket")
	}

	if _, err := s.Client.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: aws.String(s.BucketName),
	}); err != nil {
		return errors.Wrap(err, "failed to delete bucket")
	}

	return nil
}

func BenchmarkS3Save(b *testing.B) {
	s := setupS3Storage(b)
	named, ok := s.(storage.NamedStorage)
	require.True(b, ok)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		named.SaveName(context.Background(), "short", "long")
	}
}

func BenchmarkS3Load(b *testing.B) {
	s := setupS3Storage(b)
	named, ok := s.(storage.NamedStorage)
	require.True(b, ok)

	named.SaveName(context.Background(), "short", "long")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		named.Load(context.Background(), "short")
	}
}
