package reactionHandler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"lingcard-go/helpers/response"
	"lingcard-go/services/reactionService"
	"net/http"
	"strconv"
)

type ReactionHandler struct {
	reactionService *reactionService.ReactionService
}

func New(reactionService *reactionService.ReactionService) *ReactionHandler {
	return &ReactionHandler{
		reactionService: reactionService,
	}
}

func (h *ReactionHandler) Like(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "postId")
	if postId == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	id, _ := strconv.Atoi(postId)
	ctx := r.Context()
	err := h.reactionService.Like(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.CreateSuccessResponse())
}

func (h *ReactionHandler) Dislike(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "postId")
	if postId == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	id, _ := strconv.Atoi(postId)
	ctx := r.Context()
	err := h.reactionService.Dislike(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.CreateSuccessResponse())
}

func (h *ReactionHandler) Unset(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "postId")
	if postId == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	id, _ := strconv.Atoi(postId)
	ctx := r.Context()
	err := h.reactionService.Unset(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.CreateSuccessResponse())
}
