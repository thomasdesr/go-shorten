package handlers

import (
	"context"
	"html/template"
	"net/http"

	"github.com/guregu/kami"
)

type IndexParams struct {
	Short string
	Error error
}

func IndexFromContext(ctx context.Context) (IndexParams, bool) {
	p, ok := ctx.Value("IndexParams").(IndexParams)
	return p, ok
}

func IndexWithContext(ctx context.Context, ip IndexParams) context.Context {
	return context.WithValue(ctx, "IndexParams", ip)
}

func Index() kami.HandlerFunc {
	t := template.Must(template.ParseFiles("static/templates/index.tmpl"))

	return func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		params, ok := IndexFromContext(ctx)
		if !ok {
			params = IndexParams{}
		}

		err := t.Execute(w, params)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
