package server

import (
	"log"
	"net/http"
)

// RecoverMiddleware recover from panic and return error responce
func RecoverMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("recovering from %v", err)
				http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

// LoggingMiddleware log every incomming request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling", r.Method, "request for:", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
