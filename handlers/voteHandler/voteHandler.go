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
	votes, err := h.voteService.All()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{
		"success": true,
		"data":    votes,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *VoteHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	voteId := chi.URLParam(r, "id")
	if voteId == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	id, _ := strconv.Atoi(voteId)
	ctx := r.Context()
	vote, err := h.voteService.GetOne(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp := map[string]interface{}{
		"success": true,
		"data":    vote,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func (h *VoteHandler) Vote(w http.ResponseWriter, r *http.Request) {
	voteId := chi.URLParam(r, "voteOptionId")
	if voteId == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	id, _ := strconv.Atoi(voteId)
	ctx := r.Context()
	err := h.voteService.Vote(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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
	err := h.voteService.CancelVote(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.CreateSuccessResponse())
}
