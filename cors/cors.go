package cors

import (
	"net/http"
)

func WithCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	}
}
