package internal

import (
	"net/http"
)

type RequestOption func(r *http.Request)

func WithHeader(key, value string) RequestOption {
	return func(r *http.Request) {
		r.Header.Add(key, value)
	}
}

func WithQueryParam(key, value string) RequestOption {
	return func(r *http.Request) {
		q := r.URL.Query()
		q.Set(key, value)
		r.URL.RawQuery = q.Encode()
	}
}
