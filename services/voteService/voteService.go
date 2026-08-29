package voteService

import (
	"context"
	"lingcard-go/dto/voice"
	"lingcard-go/models/user"
	"lingcard-go/repositories/voiceRepository"
	"lingcard-go/repositories/voteOptionRepository"
	"lingcard-go/repositories/voteRepository"
	"time"
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
	var dto []voice.SimpleVoteDTO
	votes := s.voteRepo.All()
	for _, vote := range votes {
		dto = append(dto, voice.SimpleVoteDTO{
			ID:      vote.ID,
			Title:   vote.Title,
			Content: vote.Content,
			Voices:  s.voiceRepo.GetCountVoices(vote.ID),
		})
	}
	return dto
}

func (s *VoteService) GetOne(ctx context.Context, id int) voice.VoteDTO {
	User := ctx.Value("User").(user.User)
	Vote := s.voteRepo.Find(id)
	var voteOptionsDTO []voice.VoteOptionDTO
	var count = 0
	VoteOptions := s.voteOptionRepo.FindByVoteId(Vote.ID)
	VoteOptionsId := make([]int, len(VoteOptions))
	for i := 0; i < len(VoteOptions); i++ {
		VoteOptionsId[i] = VoteOptions[i].ID
	}
	Voice := s.voiceRepo.FindUserVoice(User.ID, VoteOptionsId)
	for _, item := range VoteOptions {
		countVoices := s.voiceRepo.CountVoices(item.ID)
		voteOptionsDTO = append(voteOptionsDTO, voice.VoteOptionDTO{
			ID:      item.ID,
			Title:   item.Title,
			Content: item.Content,
			Count:   countVoices,
		})

	}
	return voice.VoteDTO{
		ID:          Vote.ID,
		Title:       Vote.Title,
		Content:     Vote.Content,
		VoteOptions: voteOptionsDTO,
		TotalCount:  count,
		IsActive:    Vote.IsActive,
		Voted:       Voice.VoteOptionID,
	}
}

func (s *VoteService) Vote(ctx context.Context, id int) {
	User := ctx.Value("User").(user.User)
	Vote := s.voteOptionRepo.Find(id)
	VoteOptions := s.voteOptionRepo.FindByVoteId(Vote.VoteID)
	VoteOptionsId := make([]int, len(VoteOptions))
	for i := 0; i < len(VoteOptions); i++ {
		VoteOptionsId[i] = VoteOptions[i].ID
	}
	s.voiceRepo.DeleteUserVoices(User.ID, VoteOptionsId)
	var dto = voice.VoiceDTO{
		VoteOptionID: id,
		UserID:       User.ID,
		Time:         time.Now(),
	}
	s.voiceRepo.Insert(dto)
}

func (s *VoteService) CancelVote(ctx context.Context, id int) {
	User := ctx.Value("User").(user.User)
	s.voiceRepo.DeleteVoice(User.ID, id)
}
