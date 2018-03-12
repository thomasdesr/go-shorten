package main

import (
	"log"
	"net"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/jessevdk/go-flags"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/thomaso-mirodin/go-shorten/handlers"
	"github.com/thomaso-mirodin/go-shorten/storage"
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

	r := httprouter.New()

	// Serve the index
	indexPage, err := handlers.NewIndex("static/templates/index.tmpl")
	if err != nil {
		log.Fatal("Failed to create index Page", err)
	}

	r.Handler("GET", "/healthcheck", handlers.Healthcheck(store, "/healthcheck"))

	// Serve the "API"
	r.HandleMethodNotAllowed = false
	r.NotFound = handlers.GetShort(store, indexPage)
	r.Handler("POST", "/", handlers.SetShort(store))
	r.Handler("GET", "/api/search", handlers.Search(store.(storage.SearchableStorage)))

	n.UseHandler(r)

	go func() {
		log.Printf("Starting prometheus HTTP Listner on %s", net.JoinHostPort(opts.BindHost, "8081"))
		err := http.ListenAndServe(net.JoinHostPort(opts.BindHost, "8081"), promhttp.Handler())
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
