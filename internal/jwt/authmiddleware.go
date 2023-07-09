package jwt

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the request
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			jwtSecret := os.Getenv("JWT_SECRET")
			if jwtSecret == "" {
				jwtSecret = "keep it secret keep it safe"
			}

			// Verify signing method and return the secret key
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Token is valid, continue to the next handler
		next.ServeHTTP(w, r)
	})
}
