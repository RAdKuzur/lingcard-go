package wordHandler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
)

type WordHandler struct {
}

func New() *WordHandler {
	return &WordHandler{}
}

func (h *WordHandler) DownloadPackage(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	if code == "" {
		http.Error(w, "code is required", http.StatusBadRequest)
		return
	}

	filePath := "data/words/base/" + code + ".jsonl"
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "File not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	http.ServeFile(w, r, filePath)
}
