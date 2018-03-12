package handlers

import (
	"html/template"
	"log"
	"net/http"
)

var goDashboardPath = "static/templates/go-dashboard.tmpl"

func ServeGoDashboard() http.Handler {
	t, err := template.ParseFiles(goDashboardPath, searchPath)
	if err != nil {
		log.Fatal(err)
	}

	return instrumentHandler("go-dashboard", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := t.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))
}
