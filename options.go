package main

import (
	"fmt"
	"strings"

	"github.com/thomaso-mirodin/go-shorten/storage"
)

type Options struct {
	BindHost string `json:"bind_host" flag:"host" cfg:"bind_host" env:"HOST"`
	BindPort string `json:"bind_port" flag:"port" cfg:"bind_port" env:"PORT"`

	StorageType   string `json:"storage_type" flag:"storage-type" cfg:"storage_type"`
	StorageConfig string `flag:"storage-config" cfg:"storage_config"`

	// S3 Config options
	BucketName string `json:"s3_bucket" flag:"s3-bucket" cfg:"s3_bucket`
	Region     string `json:"s3_region" flag:"s3-region" cfg:"s3_region`
	AccessKey  string `json:"aws_access_key,omitempty" cfg:"aws_access_key"`
	SecretKey  string `json:"aws_secret_key,omitempty" cfg:"aws_secret_key"`

	// Inmem Config options
	InmemRandLength int `json:"inmem-length" flag:"inmem-length" cfg:"inmem-length"`

	// Filesystem Config options
	FilesystemRoot string `json:"filesystem-root", flag:"filesystem-root" cfg:"filesystem-root"`
}

func NewOptions() *Options {
	return &Options{
		StorageType:     "Inmem",
		InmemRandLength: 8,
	}
}

func (o *Options) Validate() error {
	msgs := make([]string, 0)

	if _, ok := storage.SupportedStorageTypes[o.StorageType]; !ok {
		msgs = append(msgs, fmt.Sprintf("invalid setting: storage type '%s' is not one of the supported storage types: %v", o.StorageType, storage.SupportedStorageTypes))
	}

	switch o.StorageType {
	case "Inmem":
		if o.InmemRandLength <= 0 {
			msgs = append(msgs, fmt.Sprintf("invalid setting: inmem-length is less than or equal to zero: '%d'", o.InmemRandLength))
		}
	case "S3":
		if o.BucketName == "" {
			msgs = append(msgs, "missing setting: s3-bucket")
		}
		if o.Region == "" {
			msgs = append(msgs, "missing setting: s3-region")
		}
	case "Filesystem":
		if o.FilesystemRoot == "" {
			msgs = append(msgs, "missing setting: filesystem-root")
		}
	}

	if len(msgs) != 0 {
		return fmt.Errorf("Invalid configuration:\n  %s",
			strings.Join(msgs, "\n  "))
	}
	return nil
}
