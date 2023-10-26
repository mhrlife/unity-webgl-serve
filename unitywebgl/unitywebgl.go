package unitywebgl

import (
	"net/http"
	"strings"
)

// ServeUnityWebGL serves Unity WebGL content with the appropriate headers.
func ServeUnityWebGL(folderPath string) http.Handler {
	fileServer := http.FileServer(http.Dir(folderPath))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set headers based on file extensions
		if strings.HasSuffix(r.URL.Path, ".gz") {
			w.Header().Set("Content-Encoding", "gzip")
		} else if strings.HasSuffix(r.URL.Path, ".br") {
			w.Header().Set("Content-Encoding", "br")
		}

		if strings.Contains(r.URL.Path, ".wasm") {
			w.Header().Set("Content-Type", "application/wasm")
		}

		// Serve the file
		fileServer.ServeHTTP(w, r)
	})
}
