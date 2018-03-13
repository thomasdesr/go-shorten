package handlers

import (
	"github.com/thomaso-mirodin/go-shorten/storage"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/pkg/errors"
	"log"
)

func TopN(store storage.TopN) http.Handler {
	return instrumentHandler("top_n", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n, _ := strconv.Atoi(r.URL.Query().Get("n"))
		days, _ := strconv.Atoi(r.URL.Query().Get( "days"))


		results, err := store.TopNForPeriod(r.Context(), n, days)
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

