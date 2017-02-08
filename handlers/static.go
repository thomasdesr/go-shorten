package handlers

import "net/http"

func Static(base string) http.Handler {
	sh := http.FileServer(http.Dir(base))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[1:]

		if len(path) == 0 {
			http.NotFound(w, r)
			return
		}

		sh.ServeHTTP(w, r)
	})
}
