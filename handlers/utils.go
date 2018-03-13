package handlers

import (
	"fmt"
	"net/http"
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
