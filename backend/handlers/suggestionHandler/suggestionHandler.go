package suggestionHandler

import (
	"encoding/json"
	"io"
	"lingcard-go/dto/suggestion"
	"lingcard-go/helpers/response"
	"lingcard-go/services/suggestionService"
	"net/http"
)

type SuggestionHandler struct {
	suggestionService *suggestionService.SuggestionService
}

func New(suggestionService *suggestionService.SuggestionService) *SuggestionHandler {
	return &SuggestionHandler{
		suggestionService: suggestionService,
	}
}

func (h *SuggestionHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var dto suggestion.SuggestionDTO
	err = json.Unmarshal(body, &dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx := r.Context()
	err1 := h.suggestionService.Create(ctx, dto)
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response.CreateSuccessResponse())
}
