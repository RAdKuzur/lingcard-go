package voteService

import (
	"lingcard-go/repositories/voiceRepository"
	"lingcard-go/repositories/voteOptionRepository"
	"lingcard-go/repositories/voteRepository"
)

type VoteService struct {
	voteRepo       *voteRepository.VoteRepository
	voiceRepo      *voiceRepository.VoiceRepository
	voteOptionRepo *voteOptionRepository.VoteOptionRepository
}

func New(voteRepo *voteRepository.VoteRepository, voiceRepo *voiceRepository.VoiceRepository, voteOptionRepo *voteOptionRepository.VoteOptionRepository) *VoteService {
	return &VoteService{
		voteRepo:       voteRepo,
		voiceRepo:      voiceRepo,
		voteOptionRepo: voteOptionRepo,
	}
}
