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
	"github.com/thomaso-mirodin/go-shorten/storage"
)

var opts Options

func createRouter(store storage.Storage) *kami.Mux {
	r := kami.New()

	r.Get("/*short", handlers.GetShortHandler(store))
	r.Head("/*short", handlers.GetShortHandler(store))

	r.Post("/", handlers.SetShortHandler(store))

	return r
}

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
		negroni.NewStatic(rice.MustFindBox("static").HTTPBox()),
	)

	n.UseHandler(createRouter(store))

	err = http.ListenAndServe(net.JoinHostPort(opts.BindHost, opts.BindPort), n)
	if err != nil {
		log.Fatal(err)
	}
}
