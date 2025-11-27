package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Ensure application/json header is set before any writes
		w.Header().Set("Content-Type", "application/json")

		// Use the wrapped writer to capture status and body
		wrapped := &wrappedWriter{w, http.StatusOK, bytes.Buffer{}}

		// --- REQUEST BODY LOGGING ---

		// Read the original body stream entirely
		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Warning: Failed to read request body:", err)
			// Don't interrupt flow, but proceed with empty body
		}

		// Log the body
		log.Printf("REQUEST [%s] %s\n\nBody: %s", r.Method, r.URL.Path, string(requestBody))

		// Replace the Request Body
		// CRITICAL: Give the buffered body back to the request for downstream handlers
		r.Body = io.NopCloser(bytes.NewBuffer(requestBody))

		// --- Execute Handler Chain ---
		next.ServeHTTP(wrapped, r)

		// --- RESPONSE LOGGING ---

		// Log the captured status, path, latency, and captured response body
		log.Printf("RESPONSE %d [%s] %s | Latency: %s\n\nBody: %s",
			wrapped.statusCode,
			r.Method,
			r.URL.Path,
			time.Since(start),
			wrapped.body.String(), // Access the buffered response body
		)
	})
}
