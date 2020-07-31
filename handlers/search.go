package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pkg/errors"
	"github.com/thomasdesr/go-shorten/storage"
)

func Search(store storage.SearchableStorage) http.Handler {
	return instrumentHandler("api/search", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		searchTerm := r.URL.Query().Get("s")

		results, err := store.Search(r.Context(), searchTerm)
		switch err := errors.Cause(err); err {
		case nil:
			err := json.NewEncoder(w).Encode(results)
			if err != nil {
				http.Error(w, "Failed to render JSON", http.StatusInternalServerError)
			}
		default:
			log.Printf("Error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}))
}
