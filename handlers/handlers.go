// Package handlers provides HTTP request handlers.
package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
	"github.com/thomasdesr/go-shorten/storage"
)

func Healthcheck(store storage.Storage, path string) http.Handler {
	if s, ok := store.(storage.NamedStorage); ok {
		s.SaveName(context.Background(), path, "https://google.com")
	}

	return instrumentHandler("healthcheck", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := store.Load(r.Context(), path)
		if err != nil {
			http.Error(w, "healtcheck fail", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
}

func GetShort(store storage.Storage, index Index) http.Handler {
	return instrumentHandler("get_short", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		index := Index{Template: index.Template} // Reset the index template

		short, err := getShortFromRequest(r)
		if err != nil {
			index.ServeHTTP(w, r)
			return
		}

		var shortWithoutSuffix string
		var isInfoRequest = isInfoShort(short)

		if isInfoRequest {
			shortWithoutSuffix = shortFromInfoShort(short)
		} else {
			shortWithoutSuffix = short
		}
		index.Short = shortWithoutSuffix

		url, err := store.Load(r.Context(), shortWithoutSuffix)
		switch err := errors.Cause(err); err {
		case nil:
			if isInfoRequest {
				var shortInfoResult, err = NewShortInfo()
				if err != nil {
					log.Println(err)
					index.Error = errors.Wrap(err, "I got an error :\\")
					w.WriteHeader(http.StatusInternalServerError)
				} else {
					shortInfoResult.Short = shortWithoutSuffix
					shortInfoResult.Long = url
					shortInfoResult.ServeShortInfoHTTP(w, r)
					return
				}
			} else {
				http.Redirect(w, r, url, http.StatusFound)
			}
			return
		case storage.ErrFuzzyMatchFound:
			index.Fuzzy = url
			w.WriteHeader(http.StatusNotFound)
		case storage.ErrShortNotSet:
			index.Error = fmt.Errorf("The link you specified does not exist. You can create it below.")
			w.WriteHeader(http.StatusNotFound)
		default:
			index.Error = errors.Wrap(err, "Failed to retrieve link from backend")
			w.WriteHeader(http.StatusInternalServerError)
		}

		index.ServeHTTP(w, r)
	}))
}

func infoShortResult(short string) {

}

func SetShort(store storage.NamedStorage) http.Handler {
	return instrumentHandler("set_short", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		short, err := getShortFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if short == "" {
			http.Error(w, "Missing short name", http.StatusBadRequest)
			return
		}

		url, err := getURLFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = store.SaveName(r.Context(), short, url)
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
	}))
}
