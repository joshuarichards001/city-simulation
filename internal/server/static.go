package server

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// StaticHandler manages static files and serves them.
type StaticHandler struct {
	distPath string
	fs       http.Handler
}

func NewStaticHandler(distPath string) *StaticHandler {
	return &StaticHandler{
		distPath: distPath,
		fs:       http.FileServer(http.Dir(distPath)),
	}
}

func (handler *StaticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("incoming request: method=%s path=%s remote_addr=%s",
		r.Method,
		r.URL.Path,
		r.RemoteAddr,
	)

	path := filepath.Join(handler.distPath, r.URL.Path)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		log.Printf("file not found, serving index.html: path=%s", r.URL.Path)
		http.ServeFile(w, r, filepath.Join(handler.distPath, "index.html"))
		return
	}

	fileInfo, err := os.Stat(path)
	if err == nil && fileInfo.IsDir() {
		indexPath := filepath.Join(path, "index.html")
		if _, err := os.Stat(indexPath); err == nil {
			http.ServeFile(w, r, indexPath)
			return
		}
	}

	handler.fs.ServeHTTP(w, r)
}
