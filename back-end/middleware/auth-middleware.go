package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func AdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := r.Context().Value(UserContextKey).(jwt.MapClaims)
		role := claims["role"].(string)

		if role != "admin" {
			http.Error(
				w,
				"forbidden",
				http.StatusForbidden,
			)
			return
		}

		next.ServeHTTP(w, r)

	})

}