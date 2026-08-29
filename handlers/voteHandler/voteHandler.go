package voteHandler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"lingcard-go/helpers/response"
	"lingcard-go/services/voteService"
	"net/http"
	"strconv"
)

type VoteHandler struct {
	voteService *voteService.VoteService
}

func New(voteService *voteService.VoteService) *VoteHandler {
	return &VoteHandler{
		voteService: voteService,
	}
}

func (h *VoteHandler) GetAllVotes(w http.ResponseWriter, r *http.Request) {
	votes := h.voteService.All()
	Response := map[string]interface{}{
		"success": true,
		"data":    votes,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response)
}

func (h *VoteHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	voteId := chi.URLParam(r, "id")
	if voteId == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	id, _ := strconv.Atoi(voteId)

	ctx := r.Context()

	vote := h.voteService.GetOne(ctx, id)
	Response := map[string]interface{}{
		"success": true,
		"data":    vote,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response)
}

func (h *VoteHandler) Vote(w http.ResponseWriter, r *http.Request) {
	voteId := chi.URLParam(r, "voteOptionId")
	if voteId == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	id, _ := strconv.Atoi(voteId)
	ctx := r.Context()
	h.voteService.Vote(ctx, id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.CreateSuccessResponse())
}

func (h *VoteHandler) CancelVote(w http.ResponseWriter, r *http.Request) {
	voteId := chi.URLParam(r, "voteOptionId")
	if voteId == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	id, _ := strconv.Atoi(voteId)
	ctx := r.Context()
	h.voteService.CancelVote(ctx, id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.CreateSuccessResponse())
}
