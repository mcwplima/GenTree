package middleware

import (
	"context"
	locallog "log"
	"net/http"

	"gentree/config"
	"gentree/log"

	"github.com/gorilla/handlers"
)

//WrapHandler inject
func WrapHandler(ctx context.Context, h http.Handler) http.Handler {
	l, ok := log.FromContext(ctx)
	if !ok {
		locallog.Fatal("Unable to retrieve default logger!")
	}
	config, ok := config.FromContext(ctx)
	if !ok {
		l.Fatal("Unable to retrieve default config!")
	}
	writer, err := log.Writer(config)
	if err != nil {
		l.Fatal("Unable to retrieve log writer")
	}

	h = handlers.LoggingHandler(writer, h)
	h = setCORS(h)
	h = handlers.ProxyHeaders(h)
	h = ContextHandler(ctx, h)
	return h
}

func setCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, GET, DELETE, OPTIONS, HEAD")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, Origin, Vary")
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
