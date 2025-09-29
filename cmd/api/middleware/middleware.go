package middleware

import (
	"net/http"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

type Middleware func(http.Handler) http.Handler

func CreateStack(middlewares ...Middleware) Middleware {
	// Return a function that builds and returns the full http.Handler chain
	return func(next http.Handler) http.Handler {

		// 1. Iterate BACKWARDS over the slice of middlewares
		for i := len(middlewares) - 1; i >= 0; i-- {
			middleware := middlewares[i]

			// 2. Build the chain: The result of the previous iteration
			//    becomes the 'next' handler for the current middleware.
			next = middleware(next)
		}

		// The final 'next' is the complete, correctly ordered chain
		return next
	}
}

// func CreateStack(middlewares ...Middleware) Middleware {
// 	return func(next http.Handler) http.Handler {
// 		for _, middleware := range middlewares {
// 			next = middleware(next)
// 		}
// 		return next
// 	}
// }
