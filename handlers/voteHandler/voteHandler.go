package voteHandler

import (
	"encoding/json"
	"lingcard-go/services/voteService"
	"net/http"
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(votes)
}

func (h *VoteHandler) GetOne(w http.ResponseWriter, r *http.Request) {

}

func (h *VoteHandler) Vote(w http.ResponseWriter, r *http.Request) {

}

func (h *VoteHandler) CancelVote(w http.ResponseWriter, r *http.Request) {

}
