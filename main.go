package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

const (
	defaultPort      = "8080"
	frontendDistPath = "./frontend/dist"
)

func main() {
	// Initialize structured logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Create file server for static assets
	fs := http.FileServer(http.Dir(frontendDistPath))

	// Serve static files with SPA fallback
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Log incoming request
		slog.Info("incoming request",
			"method", r.Method,
			"path", r.URL.Path,
			"remote_addr", r.RemoteAddr,
		)

		// Build the full file path
		path := filepath.Join(frontendDistPath, r.URL.Path)

		// Check if file exists
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			// If file doesn't exist, serve index.html (SPA fallback)
			slog.Debug("file not found, serving index.html", "path", r.URL.Path)
			http.ServeFile(w, r, filepath.Join(frontendDistPath, "index.html"))
			return
		}

		// If it's a directory, try to serve index.html from that directory
		fileInfo, err := os.Stat(path)
		if err == nil && fileInfo.IsDir() {
			indexPath := filepath.Join(path, "index.html")
			if _, err := os.Stat(indexPath); err == nil {
				http.ServeFile(w, r, indexPath)
				return
			}
		}

		// Serve the requested file
		fs.ServeHTTP(w, r)
	})

	// Start server
	addr := ":" + port
	slog.Info("starting server", "port", port, "static_path", frontendDistPath)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("server failed to start", "error", err)
	}
}
