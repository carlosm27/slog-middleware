package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	port := ":8000"

	logger := slog.Default()
	finalHandler := http.HandlerFunc(homeHandler)
	mux.Handle("/", LoggingMiddleware(finalHandler))

	logger.Info("Listening on ", "port", port)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		logger.Error("Problem starting the server", "error", err)
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("This is my home page"))

}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//var resp http.Response

		logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
		logger.Info("Request Info",
			slog.String("method", r.Method),
			slog.String("path", r.RequestURI),
			//slog.String("status", StatusCode),
		)
		next.ServeHTTP(w, r)
		statusCode := w.Header().Get("Status")
		logger.Info("Request Info",
			slog.String("method", r.Method),
			slog.String("path", r.RequestURI),
			slog.String("status", statusCode),
		)
	})
}
