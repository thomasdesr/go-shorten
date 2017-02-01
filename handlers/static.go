package handlers

import (
	"context"
	"net/http"

	"github.com/guregu/kami"
)

func Static(base string) kami.HandlerType {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		path := kami.Param(ctx, "path")

		if len(path) == 0 {
			http.NotFound(w, r)
			return
		}

		http.FileServer(http.Dir(base)).ServeHTTP(w, r)
	}
}
