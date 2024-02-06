package middleware

import (
	"log"
	"net/http"
	"strings"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestURI := r.RequestURI
		if strings.HasSuffix(requestURI, "/") {
			requestURI = requestURI[:len(requestURI)-1]
		}
		log.Println("Got request:", requestURI)
		next.ServeHTTP(w, r)
	},
	)
}
