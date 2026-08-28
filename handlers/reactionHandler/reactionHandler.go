package reactionHandler

import (
	"github.com/go-chi/chi/v5"
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
	h.reactionService.Like(ctx, id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *ReactionHandler) Dislike(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "postId")
	if postId == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	id, _ := strconv.Atoi(postId)
	ctx := r.Context()
	h.reactionService.Dislike(ctx, id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *ReactionHandler) Unset(w http.ResponseWriter, r *http.Request) {
	postId := chi.URLParam(r, "postId")
	if postId == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	id, _ := strconv.Atoi(postId)
	ctx := r.Context()
	h.reactionService.Unset(ctx, id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
