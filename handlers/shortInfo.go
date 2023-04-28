package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type ShortInfo struct {
	Short    string
	Long     string
	Error    string
	Template *template.Template
}

var shortInfoTemplatePath = "static/templates/existing.tmpl"

func NewShortInfo() (ShortInfo, error) {
	t, err := template.ParseFiles(shortInfoTemplatePath)
	if err != nil {
		return ShortInfo{}, err
	}

	return ShortInfo{Template: t}, nil
}

func (i ShortInfo) ServeShortInfoHTTP(w http.ResponseWriter, r *http.Request) {
	err := i.Template.Execute(w, i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}
