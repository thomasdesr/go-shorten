package main

import (
	"log"
	"net"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/codegangsta/negroni"
	"github.com/jessevdk/go-flags"
	"github.com/julienschmidt/httprouter"
	"github.com/thomaso-mirodin/go-shorten/handlers"
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
		negroni.NewStatic(rice.MustFindBox("static").HTTPBox()),
	)

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
