package middleware

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Key to store user information in context
type contextKey string

const UserContextKey contextKey = "user"

// Generate JWT Token
func GenerateToken(email string, isAdmin bool, id int64) (string, error) {
	jwtSecret := strings.TrimSpace(os.Getenv("SECRET_KEY"))

	claims := jwt.MapClaims{
		"user_id":  id,
		"email":    email,
		"is_admin": isAdmin,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(), // Token valid for 30 day
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// Function to get user data from request
func GetUserFromContext(r *http.Request) (map[string]any, bool) {
	userData, ok := r.Context().Value(UserContextKey).(map[string]any)
	return userData, ok
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

		// Extract claims from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid Token Claims", http.StatusUnauthorized)
			return
		}

		// Extract data from token
		userId, _ := claims["user_id"].(float64) // JWT stores numbers as float64
		email, _ := claims["email"].(string)
		isAdmin, _ := claims["is_admin"].(bool)

		// Store the user data in the request context
		ctx := context.WithValue(r.Context(), UserContextKey, map[string]any{
			"user_id":  int64(userId),
			"email":    email,
			"is_admin": isAdmin,
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := GetUserFromContext(r)

		if !ok {
			http.Error(w, "Unauthorized, please re login!", http.StatusUnauthorized)
			return
		}

		isAdmin := user["is_admin"].(bool)

		if !isAdmin {
			http.Error(w, "You are not authorized to perform this action!", http.StatusForbidden)
			return
		}

		next(w, r)
	}
}
