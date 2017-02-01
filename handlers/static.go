package handlers

import (
	"context"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/guregu/kami"
)

func Static(box *rice.Box) kami.Middleware {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) context.Context {
		name := r.URL.Path

		f, err := box.Open(name)
		if err != nil {
			// If we don't find the file in the Box, pass the request through
			return ctx
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			// If we can't get metadata about the file the file, pass the request through
			return ctx
		}

		if fi.IsDir() {
			// If directory, pass the request through
			return ctx
		}

		// If we have the file, serve it and don't pass the request
		http.ServeContent(w, r, name, fi.ModTime(), f)
		return nil
	}
}
