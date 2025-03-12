package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Generate JWT Token
func GenerateToken(email string, id int64) (string, error) {
	jwtSecret := strings.TrimSpace(os.Getenv("SECRET_KEY"))

	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(), // Token valid for 30 day
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// Auth Middleware
func Authentication(next http.Handler) http.Handler {
	jwtSecret := os.Getenv("SECRET_KEY")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
