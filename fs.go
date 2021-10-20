package kgo

import (
	"net/http"
	"path/filepath"
)

// HTTPFS returns the index.html of directory instead of file list
type HTTPFS struct {
	FS http.FileSystem
}

// Open ...
func (fs HTTPFS) Open(path string) (http.File, error) {
	f, err := fs.FS.Open(path)
	if err != nil {
		return nil, err
	}
	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := fs.FS.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
