package handlers

import (
	"html/template"
	"net/http"
)

type Index struct {
	Short    string
	Error    error
	template *template.Template
}

var defaultIndexPath = "static/templates/index.tmpl"

func NewIndex(path string) (Index, error) {
	t, err := template.ParseFiles(path)
	if err != nil {
		return Index{}, err
	}

	return Index{template: t}, nil
}

func (i Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := i.template.Execute(w, i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
