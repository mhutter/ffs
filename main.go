package ffs

import (
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// New returns a handler that serves static assets from root, but serves
// "index.html" from root if the requested file does not exist.
func New(root string) *http.ServeMux {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(root))

	for _, path := range findFiles(root) {
		mux.Handle(path, fs)
	}

	// Fallback: send index.html
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Join(root, "index.html"))
	})

	return mux
}

func findFiles(root string) (files []string) {
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() {
			files = append(files, strings.TrimPrefix(path, root))
		}
		return nil
	})

	if err != nil {
		log.Fatalln("ERROR scanning static asset dir:", err)
	}

	return
}
