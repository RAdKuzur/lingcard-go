package voteService

import (
	"lingcard-go/dto/voice"
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

func (s *VoteService) All() []voice.SimpleVoteDTO {
	votes := s.voteRepo.All()
	for _, vote := range votes {
		vote.Voices = s.voiceRepo.GetCountVoices(vote.ID)
	}
	return votes
}
