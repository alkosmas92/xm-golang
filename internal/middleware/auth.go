package middleware

import (
	"context"
	"net/http"
	"strings"

	appContext "github.com/alkosmas92/xm-golang/internal/context"
	"github.com/alkosmas92/xm-golang/internal/utils"
)

// AuthMiddleware handles authentication for HTTP requests.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token := parts[1]
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if claims.UserID == "" {
			http.Error(w, "invalid token: empty userID", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), appContext.UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, appContext.UsernameKey, claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
