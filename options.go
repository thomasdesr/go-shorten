package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/goamz/goamz/aws"
	"github.com/thomaso-mirodin/go-shorten/storage"
)

type Options struct {
	BindHost string `long:"host"    default:"0.0.0.0"   env:"HOST"`
	BindPort string `long:"port"    default:"8080"      env:"PORT"`

	StorageType string `long:"storage-type" default:"Inmem" ini-name:"storage_type" env:"STORAGE_TYPE"`
	// StorageConfig string `long:"storage-config" ini-name:"storage_config"`

	// S3 Config options
	S3 struct {
		BucketName string `long:"s3-bucket" default:"go-shorten"    env:"S3_BUCKET"`
		Region     string `long:"s3-region" default:"us-west-2"     env:"S3_REGION"`
		AccessKey  string `                                         env:"AWS_ACCESS_KEY"`
		SecretKey  string `                                         env:"AWS_SECRET_KEY"`
	} `group:"S3 Storage Options"`

	// Inmem Config options
	Inmem struct {
		RandLength int `long:"inmem-length" default:"8" env:"INMEM_LENGTH"`
	} `group:"Inmem Storage Options"`

	// Filesystem Config options
	Filesystem struct {
		RootPath string `long:"root-path" default:"./url-storage" env:"ROOT_PATH"`
	} `group:"Filesystem Storage Options"`
}

// createStorageFromOption takes an Option struct and based on the StorageType field constructs a storage.Storage and returns it.
func createStorageFromOption(opts *Options) (storage.Storage, error) {
	switch opts.StorageType {
	case "Inmem":
		log.Printf("Setting up an Inmem Storage layer with short code length of '%d'", opts.Inmem.RandLength)

		return storage.NewInmem(opts.Inmem.RandLength)
	case "S3":
		log.Println("Setting up an S3 Storage layer")
		log.Printf("Connecting to AWS Region '%s' for bucket '%s'", opts.S3.Region, opts.S3.BucketName)

		region, ok := aws.Regions[opts.S3.Region]
		if !ok {
			log.Fatalf("Unable to find a region that matches '%s'", opts.S3.Region)
		}

		if len(opts.S3.BucketName) == 0 {
			log.Fatalf("BucketName has be something (currently has zero length)")
		}

		auth, err := aws.GetAuth(opts.S3.AccessKey, opts.S3.SecretKey, "", time.Time{})
		if err != nil {
			akl := len(opts.S3.AccessKey)
			skl := len(opts.S3.SecretKey)
			log.Printf("Credential lengths: accessKey(%v) secretKey(%v)", akl, skl)
			log.Fatalf("Unable to find valid auth credentials because: %v", err)
		}

		return storage.NewS3(auth, region, opts.S3.BucketName)
	case "Filesystem":
		log.Println("Setting up a Filesystem storag layer with root: %v", opts.Filesystem.RootPath)

		return storage.NewFilesystem(opts.Filesystem.RootPath)
	default:
		validTypes := make([]string, len(storage.SupportedStorageTypes))

		i := 0
		for k := range storage.SupportedStorageTypes {
			validTypes[i] = k
			i++
		}

		return nil, fmt.Errorf("Unsupported storage-type: '%s', valid ones are: '%v'", opts.StorageType, strings.Join(validTypes, "', '"))
	}
}
