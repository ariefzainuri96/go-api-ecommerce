package middleware

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Skip Swagger & static assets
		swaggerPrefixes := []string{"/swagger", "/docs", "/swagger-ui", "/openapi"}
		for _, p := range swaggerPrefixes {
			if strings.HasPrefix(r.URL.Path, p) {
				next.ServeHTTP(w, r)
				return
			}
		}

		// Limit size
		const maxSize = 2048

		// Ensure application/json header is set before any writes
		w.Header().Set("Content-Type", "application/json")

		// Use the wrapped writer to capture status and body
		wrapped := &wrappedWriter{w, http.StatusOK, bytes.Buffer{}}

		if r.Method != http.MethodGet {
			// --- REQUEST BODY LOGGING ---

			// Read the original body stream entirely
			requestBody, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println("Warning: Failed to read request body:", err)
				// Don't interrupt flow, but proceed with empty body
			}

			if len(requestBody) > maxSize {
				requestBody = requestBody[:maxSize]
			}

			// Log the body
			log.Printf("REQUEST [%s] %s\n\nBody: %s", r.Method, r.URL.Path, string(requestBody))

			// Replace the Request Body
			// CRITICAL: Give the buffered body back to the request for downstream handlers
			r.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		// --- Execute Handler Chain ---
		next.ServeHTTP(wrapped, r)

		// --- RESPONSE LOGGING ---

		var query strings.Builder

		first := true
		for key, values := range r.URL.Query() {
			for _, v := range values {				
				if !first {
					query.WriteString("&")
				} else {
					query.WriteString("?")
				}
				first = false
				query.WriteString(fmt.Sprintf("%s=%s", key, v))
			}
		}

		// Log the captured status, path, latency, and captured response body
		log.Printf("RESPONSE %d [%s] %s | Latency: %s\n\nBody: %s",
			wrapped.statusCode,
			r.Method,
			fmt.Sprintf("%v%v", r.URL.Path, query.String()),
			time.Since(start),
			wrapped.body.String(), // Access the buffered response body
		)
	})
}
