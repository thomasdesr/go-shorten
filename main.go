package main

import (
	"log"
	"net"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/codegangsta/negroni"
	"github.com/guregu/kami"
	"github.com/jessevdk/go-flags"
	"github.com/thomaso-mirodin/go-shorten/handlers"
	"github.com/thomaso-mirodin/go-shorten/handlers/templates"
)

var opts Options

//go:generate rice embed-go -v

func main() {
	if _, err := flags.Parse(&opts); err != nil {
		return
	}

	store, err := createStorageFromOption(&opts)
	if err != nil {
		log.Fatal(err)
	}

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
	)

	r := kami.New()

	box := rice.MustFindBox("static")

	// Serve the static content
	r.Use("/", handlers.Static(box))

	// Serve the index
	r.Get("/", templates.Index(box))

	// Serve the "API"
	r.Get("/*short", handlers.GetShortHandler(store))
	r.Post("/", handlers.SetShortHandler(store))

	n.UseHandler(r)

	err = http.ListenAndServe(net.JoinHostPort(opts.BindHost, opts.BindPort), n)
	if err != nil {
		log.Fatal(err)
	}
}
