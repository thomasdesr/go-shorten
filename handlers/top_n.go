package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/thomasdesr/go-shorten/storage"
)

func TopN(store storage.TopN) http.Handler {
	return instrumentHandler("top_n", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n, _ := strconv.Atoi(r.URL.Query().Get("n"))
		days, _ := strconv.Atoi(r.URL.Query().Get("days"))

		results, err := store.TopNForPeriod(r.Context(), n, days)
		if err != nil {
			log.Printf("TopN error: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(results)
		if err != nil {
			http.Error(w, "Failed to render JSON", http.StatusInternalServerError)
		}
	}))
}
