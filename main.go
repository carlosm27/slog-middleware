package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	port := ":8080"

	// Register the routes and handlers
	finalHandler := http.HandlerFunc(homeHandler)
	mux.Handle("/", LoggingMiddleware(finalHandler))

	slog.Info("Listening on ", "port", port)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		slog.Warn("Problem starting the server", "error", err)
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("This is my home page"))

}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
		logger.Info("Request Info",
			slog.String("method", r.Method),
			slog.String("path", r.RequestURI),
		)
		next.ServeHTTP(w, r)
	})
}
