package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/goamz/aws"
	"github.com/mreiferson/go-options"
	"github.com/thomaso-mirodin/go-shorten/handlers"
	"github.com/thomaso-mirodin/go-shorten/storage"
)

func parseArgsIntoOptions() *Options {
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	config := flagSet.String("config", "", "path to config file")

	flagSet.String("host", "0.0.0.0", "Which interface to bind to?")
	flagSet.String("port", "8080", "Which port to bind to?")

	flagSet.String("storage-type", "Inmem", "Which storage engine to use")
	flagSet.String("storage-config", "", "Path to storage config")

	// S3 Arguments
	flagSet.String("s3-bucket", "go-shorten", "Name of the S3 bucket to use for storage")
	flagSet.String("s3-region", "us-west-2", "AWS region to connect to")

	// Inmem Arguments
	flagSet.Int("inmem-length", 8, "How long should the short url be for inmem storage")

	// Filesystem Arguments
	flagSet.String("filesystem-root", "./url-storage", "Directory where the filesystem storage will store its files")

	flagSet.Parse(os.Args[1:])

	cfg := make(EnvOptions)
	if *config != "" {
		f, err := os.Open(*config)
		if err != nil {
			log.Fatalf("Failed to open config file %s: %s", *config, err)
		}
		if err := json.NewDecoder(f).Decode(&cfg); err != nil {
			log.Fatalf("Failed to load config file %s - %s", *config, err)
		}
	}

	opts := NewOptions()
	cfg.LoadEnvForStruct(opts)
	options.Resolve(opts, flagSet, cfg)

	return opts
}

// createStorageFromOption takes an Option struct and based on the StorageType field constructs a storage.Storage and returns it.
func createStorageFromOption(opts *Options) (storage.Storage, error) {
	switch opts.StorageType {
	case "Inmem":
		log.Printf("Setting up an Inmem Storage layer with short code length of '%d'", opts.InmemRandLength)

		return storage.NewInmem(opts.InmemRandLength)
	case "S3":
		log.Println("Setting up an S3 Storage layer")
		log.Printf("Connecting to AWS Region '%s' for bucket '%s'", opts.Region, opts.BucketName)

		auth, err := aws.GetAuth(opts.AccessKey, opts.SecretKey)
		if err != nil {
			log.Fatal("Unable to find valid auth credentials because: %v", err)
		}

		region, ok := aws.Regions[opts.Region]
		if !ok {
			log.Fatalf("Unable to find a region that matches '%s'", opts.Region)
		}

		return storage.NewS3(auth, region, opts.BucketName)
	case "Filesystem":
		log.Println("Setting up a Filesystem storag layer with root: %v", opts.FilesystemRoot)

		return storage.NewFilesystem(opts.FilesystemRoot)
	default:
		return nil, fmt.Errorf("Unsupported storage-type: '%s'", opts.StorageType)
	}
}

func main() {
	opts := parseArgsIntoOptions()
	if err := opts.Validate(); err != nil {
		log.Fatal(err)
	}

	store, err := createStorageFromOption(opts)
	if err != nil {
		log.Fatal(err)
	}

	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger(), negroni.NewStatic(http.Dir("static")))

	r := httprouter.New()

	r.GET("/*short", handlers.GetShortHandler(store))
	r.HEAD("/*short", handlers.GetShortHandler(store))

	r.POST("/*short", handlers.SetShortHandler(store))
	r.PUT("/*short", handlers.SetShortHandler(store))

	n.UseHandler(r)

	err = http.ListenAndServe(net.JoinHostPort(opts.BindHost, opts.BindPort), n)
	if err != nil {
		log.Fatal(err)
	}
}
