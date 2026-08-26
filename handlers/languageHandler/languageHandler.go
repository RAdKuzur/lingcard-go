package languageHandler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"lingcard-go/services/languageService"
	"net/http"
	"strconv"
)

type LanguageHandler struct {
	langService *languageService.LanguageService
}

func New(langService *languageService.LanguageService) *LanguageHandler {
	return &LanguageHandler{langService: langService}
}

func (h *LanguageHandler) GetAllLanguages(w http.ResponseWriter, r *http.Request) {
	var languages = h.langService.All()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(languages)
}

func (h *LanguageHandler) GetAllActiveLanguages(w http.ResponseWriter, r *http.Request) {
	var languages = h.langService.AllActive()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(languages)
}

func (h *LanguageHandler) GetExceptLanguages(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	languageID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	var languages = h.langService.ExceptLanguage(languageID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(languages)
}

func (h *LanguageHandler) GetLanguageMap(w http.ResponseWriter, r *http.Request) {
	var languages = h.langService.Map()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(languages)
}
