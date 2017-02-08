package main

import (
	"log"
	"net"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/guregu/kami"
	"github.com/jessevdk/go-flags"
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
	)

	r := kami.New()

	// Serve the static content
	static := handlers.Static("static")
	r.Get("/css/*path", static)
	r.Get("/js/*path", static)
	r.Get("/img/*path", static)

	// Serve the index
	indexPage, err := handlers.NewIndex("static/templates/index.tmpl")
	if err != nil {
		log.Fatal("Failed to create index Page", err)
	}
	r.Get("/", indexPage)

	// Serve the "API"
	r.Get("/*short", handlers.GetShortHandler(store, indexPage))
	r.Post("/", handlers.SetShortHandler(store))

	n.UseHandler(r)

	log.Printf("Starting HTTP Listener on: %s", net.JoinHostPort(opts.BindHost, opts.BindPort))
	err = http.ListenAndServe(net.JoinHostPort(opts.BindHost, opts.BindPort), n)
	if err != nil {
		log.Fatal(err)
	}
}
