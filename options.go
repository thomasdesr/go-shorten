package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/shlex"
	flags "github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
	"github.com/thomaso-mirodin/go-shorten/storage"
	"github.com/thomaso-mirodin/go-shorten/storage/multistorage"
)

type Options struct {
	BindHost string `long:"host"    default:"0.0.0.0"   env:"HOST"`
	BindPort string `long:"port"    default:"8080"      env:"PORT"`

	StorageType string `long:"storage-type" default:"Inmem" ini-name:"storage_type" env:"STORAGE_TYPE"`
	// StorageConfig string `long:"storage-config" ini-name:"storage_config"`

	// S3 Config options
	S3 struct {
		BucketName string `long:"s3-bucket" default:"go-shorten"    env:"S3_BUCKET"`
	} `group:"S3 Storage Options"`

	// Inmem Config options
	Inmem struct {
		RandLength int `long:"inmem-length" default:"8" env:"INMEM_LENGTH"`
	} `group:"Inmem Storage Options"`

	// Filesystem Config options
	Filesystem struct {
		RootPath string `long:"root-path" default:"./url-storage" env:"ROOT_PATH"`
	} `group:"Filesystem Storage Options"`

	Regex struct {
		Remaps map[string]string `long:"regex-remap" env:"REGEX_REMAP"`
	} `group:"Regex Storage Options"`

	Multistorage struct {
		StorageArgs []string `long:"multi-sub-args" env:"MULTI_SUB_ARGS"`
	} `group:"Multi Storage Options"`
}

// createStorageFromOption takes an Option struct and based on the StorageType field constructs a storage.Storage and returns it.
func createStorageFromOption(opts *Options) (storage.Storage, error) {
	switch opts.StorageType {
	case "Inmem":
		log.Printf("Setting up an Inmem Storage layer with short code length of '%d'", opts.Inmem.RandLength)

		return storage.NewInmem(opts.Inmem.RandLength)
	case "S3":
		log.Println("Setting up an S3 Storage layer")

		if len(opts.S3.BucketName) == 0 {
			log.Fatalf("BucketName has be something (currently empty)")
		}

		return storage.NewS3(nil, opts.S3.BucketName)
	case "Filesystem":
		log.Printf("Setting up a Filesystem storage layer with root: %v", opts.Filesystem.RootPath)

		return storage.NewFilesystem(opts.Filesystem.RootPath)
	case "Regex":
		log.Printf("Setting up a Regex storage with %v remaps", opts.Regex.Remaps)

		return storage.NewRegexFromList(opts.Regex.Remaps)
	case "Multistorage":
		log.Printf("Setting up a Multilayer Storage layer with children: %q", opts)
		log.Println("WARNING: Multilayer currently only works in a READ ONLY mode")

		storages := make([]storage.NamedStorage, 0, len(opts.Multistorage.StorageArgs))
		for i, rawArgs := range opts.Multistorage.StorageArgs {
			subArgs, err := shlex.Split(rawArgs)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to split arguments for argument #%d", i)
			}

			var subOpt Options
			_, err = flags.ParseArgs(&subOpt, subArgs)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to cli parse sub argument #%d", i)
			}

			store, err := createStorageFromOption(&subOpt)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to create storage #%d from args", i)
			}

			nstore, ok := store.(storage.NamedStorage)
			if !ok {
				return nil, errors.New("MultiStorage only supports NamedStorage backends")
			}

			storages = append(storages, nstore)
		}

		return multistorage.Simple(storages...)
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
