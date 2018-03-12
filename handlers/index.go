package handlers

import (
	"html/template"
	"net/http"
)

type Index struct {
	Short    string
	Error    error
	Fuzzy    string
	Template *template.Template
}

var defaultIndexPath = "static/templates/index.tmpl"
var searchPath = "static/templates/search.tmpl"

func NewIndex(path string) (Index, error) {
	t, err := template.ParseFiles(path, searchPath)
	if err != nil {
		return Index{}, err
	}

	return Index{Template: t}, nil
}

func (i Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := i.Template.Execute(w, i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
