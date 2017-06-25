// Package handlers provides HTTP request handlers.
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/thomaso-mirodin/go-shorten/storage"
)

func getShortFromRequest(r *http.Request) (short string, err error) {
	if short := r.URL.Path[1:]; len(short) > 0 {
		return short, nil
	}

	if short := r.PostFormValue("code"); len(short) > 0 {
		return short, nil
	}

	return "", fmt.Errorf("failed to find short in request")
}

func getURLFromRequest(r *http.Request) (url string, err error) {
	if url := r.PostFormValue("url"); len(url) > 0 {
		return url, nil
	}

	return "", fmt.Errorf("failed to find short in request")
}

func GetShortHandler(store storage.Storage, index Index) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		index := Index{Template: index.Template} // Reset the index template

		short, err := getShortFromRequest(r)
		if err != nil {
			index.ServeHTTP(w, r)
			return
		}
		index.Short = short

		url, err := store.Load(r.Context(), short)
		switch err := errors.Cause(err); err {
		case nil:
			http.Redirect(w, r, url, http.StatusFound)
			return
		case storage.ErrShortNotSet:
			index.Error = fmt.Errorf("The link you specified does not exist. You can create it below.")
			w.WriteHeader(http.StatusNotFound)
		default:
			index.Error = errors.Wrap(err, "Failed to retrieve link from backend")
			w.WriteHeader(http.StatusInternalServerError)
		}

		index.ServeHTTP(w, r)
	})
}

func SetShortHandler(store storage.Storage) http.Handler {
	named, namedOk := store.(storage.NamedStorage)
	unnamed, unnamedOk := store.(storage.UnnamedStorage)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		short, err := getShortFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		url, err := getURLFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if short == "" {
			if !unnamedOk {
				http.Error(w, "Current storage layer does not support storing an unnamed url", http.StatusBadRequest)
				return
			}

			short, err = unnamed.Save(r.Context(), url)
		} else {
			if !namedOk {
				http.Error(w, "Current storage layer does not support storing a named url", http.StatusBadRequest)
				return
			}

			err = named.SaveName(r.Context(), short, url)
		}
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to save '%s' to '%s' because: %s", url, short, err), http.StatusInternalServerError)
			return
		}

		// Return the short code formatted based on Accept headers
		switch r.Header.Get("Accept") {
		case "application/json":
			w.Header().Set("Content-Type", "application/json; charset=utf-8")

			err := json.NewEncoder(w).Encode(map[string]string{"short": short, "url": url})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "text/plain":
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			fmt.Fprintln(w, short)
		default:
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			fmt.Fprintln(w, short)
		}
	})
}
