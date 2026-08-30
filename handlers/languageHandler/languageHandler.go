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
	var languages, err = h.langService.All()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{
		"success": true,
		"data":    languages,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *LanguageHandler) GetAllActiveLanguages(w http.ResponseWriter, r *http.Request) {
	var languages, err = h.langService.AllActive()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{
		"success": true,
		"data":    languages,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *LanguageHandler) GetExceptLanguages(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	languageID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	languages, err2 := h.langService.ExceptLanguage(languageID)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{
		"success": true,
		"data":    languages,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *LanguageHandler) GetLanguageMap(w http.ResponseWriter, r *http.Request) {
	var languages, err = h.langService.Map()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"map": languages,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
