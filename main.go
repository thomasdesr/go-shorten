package main

import (
	"log"
	"net"
	"net/http"

	"github.com/jessevdk/go-flags"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/thomasdesr/go-shorten/handlers"
	"github.com/thomasdesr/go-shorten/storage"
	"github.com/urfave/negroni"
)

var opts Options

func main() {
	if _, err := flags.Parse(&opts); err != nil {
		return
	}

	store, err := createStorageFromOption(&opts)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Storage successfully created")

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir("static")),
	)

	mux := http.NewServeMux()
	mux.Handle("GET /healthcheck", handlers.Healthcheck(store, "/healthcheck"))

	// Serve the index
	indexPage, err := handlers.NewIndex("static/templates/index.tmpl")
	if err != nil {
		log.Fatal("Failed to create index Page", err)
	}

	// Go Endpoints
	mux.Handle("GET /go", handlers.ServeGoDashboard())

	// API handlers
	mux.Handle("POST /", handlers.SetShort(store)) // TODO(@thomas): move this to a stable API endpoint
	if ss, ok := store.(storage.SearchableStorage); ok {
		mux.Handle("GET /_api/v1/search", handlers.Search(ss))
	}
	if tns, ok := store.(storage.TopN); ok {
		mux.Handle("GET /_api/v1/top_n", handlers.TopN(tns))
	}

	// Catch-all handler for short URLs (must be last)
	mux.Handle("/", handlers.GetShort(store, indexPage))

	n.UseHandler(mux)

	go func() {
		log.Printf("Starting prometheus HTTP Listner on %s", net.JoinHostPort(opts.BindHost, opts.MetricsPort))
		err := http.ListenAndServe(net.JoinHostPort(opts.BindHost, opts.MetricsPort), promhttp.Handler())
		if err != nil {
			log.Println(err)
		}
	}()

	log.Printf("Starting HTTP Listener on: %s", net.JoinHostPort(opts.BindHost, opts.BindPort))
	err = http.ListenAndServe(net.JoinHostPort(opts.BindHost, opts.BindPort), n)
	if err != nil {
		log.Fatal(err)
	}
}
