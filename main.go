package main

import (
	"log"
	"net"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/guregu/kami"
	"github.com/jessevdk/go-flags"
	"github.com/thomaso-mirodin/go-shorten/handlers"
	"github.com/thomaso-mirodin/go-shorten/handlers/templates"
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

	r.Get("/css/*path", handlers.Static("static"))
	r.Get("/js/*path", handlers.Static("static"))
	r.Get("/img/*path", handlers.Static("static"))
	// r.Get("/static", http.FileServer(http.Dir(".")))

	// Serve the index
	r.Get("/", templates.Index())

	// Serve the "API"
	r.Get("/*short", handlers.GetShortHandler(store))
	r.Post("/", handlers.SetShortHandler(store))

	n.UseHandler(r)

	log.Printf("Starting HTTP Listener on: %s", net.JoinHostPort(opts.BindHost, opts.BindPort))
	err = http.ListenAndServe(net.JoinHostPort(opts.BindHost, opts.BindPort), n)
	if err != nil {
		log.Fatal(err)
	}
}
