package middleware

import (
	"context"
	"net/http"

	"specialist/internal/db"
)

func DBMiddleware(connector *db.Connector, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "connector", connector)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
