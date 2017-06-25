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

	log.Println("Storage created:", opts)

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

	// Serve the "API"
	getHandler := handlers.GetShortHandler(store, indexPage)
	r.Handler("HEAD", "/*short", getHandler)
	r.Handler("GET", "/*short", getHandler)
	r.Handler("POST", "/", handlers.SetShortHandler(store))

	n.UseHandler(r)

	go func() {
		log.Printf("Starting prometheus HTTP Listner on %s", net.JoinHostPort(opts.BindHost, "8081"))
		err := http.ListenAndServe(net.JoinHostPort(opts.BindHost, "8081"), promhttp.Handler())
		if err != nil {
			panic(err)
		}
	}()

	log.Printf("Starting HTTP Listener on: %s", net.JoinHostPort(opts.BindHost, opts.BindPort))
	err = http.ListenAndServe(net.JoinHostPort(opts.BindHost, opts.BindPort), n)
	if err != nil {
		log.Fatal(err)
	}
}
