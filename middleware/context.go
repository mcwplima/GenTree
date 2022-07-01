package middleware

import (
	"context"
	"net/http"
)

// ContextHandler attaches an context to the request
func ContextHandler(ctx context.Context, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		*r = *r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
