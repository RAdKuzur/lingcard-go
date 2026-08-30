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

func (s *VoteService) All() ([]voice.SimpleVoteDTO, error) {
	var dto = make([]voice.SimpleVoteDTO, 0)
	votes, err := s.voteRepo.All()
	if err != nil {
		return dto, err
	}
	for _, vote := range votes {
		count, err1 := s.voiceRepo.GetCountVoices(vote.ID)
		if err1 != nil {
			return dto, err1
		}
		dto = append(dto, voice.SimpleVoteDTO{
			ID:      vote.ID,
			Title:   vote.Title,
			Content: vote.Content,
			Voices:  count,
		})
	}
	return dto, err
}

func (s *VoteService) GetOne(ctx context.Context, id int) (voice.VoteDTO, error) {
	var count = 0
	var voteOptionsDTO []voice.VoteOptionDTO
	usr := ctx.Value("User").(user.User)
	vt, err := s.voteRepo.Find(id)
	if err != nil {
		return voice.VoteDTO{}, err
	}
	vtOptions, err1 := s.voteOptionRepo.FindByVoteId(vt.ID)
	if err1 != nil {
		return voice.VoteDTO{}, err1
	}
	vtOptionsId := make([]int, len(vtOptions))
	for i := 0; i < len(vtOptions); i++ {
		vtOptionsId[i] = vtOptions[i].ID
	}
	vc, err2 := s.voiceRepo.FindUserVoice(usr.ID, vtOptionsId)
	if err2 != nil {
		return voice.VoteDTO{}, err2
	}
	for _, item := range vtOptions {
		countVoices, err3 := s.voiceRepo.CountVoices(item.ID)
		if err3 != nil {
			return voice.VoteDTO{}, err3
		}
		voteOptionsDTO = append(voteOptionsDTO, voice.VoteOptionDTO{
			ID:      item.ID,
			Title:   item.Title,
			Content: item.Content,
			Count:   countVoices,
		})

	}
	return voice.VoteDTO{
		ID:          vt.ID,
		Title:       vt.Title,
		Content:     vt.Content,
		VoteOptions: voteOptionsDTO,
		TotalCount:  count,
		IsActive:    vt.IsActive,
		Voted:       vc.VoteOptionID,
	}, err
}

func (s *VoteService) Vote(ctx context.Context, id int) error {
	usr := ctx.Value("User").(user.User)
	vt, err := s.voteOptionRepo.Find(id)
	if err != nil {
		return err
	}
	vtOptions, err1 := s.voteOptionRepo.FindByVoteId(vt.VoteID)
	if err1 != nil {
		return err1
	}
	VoteOptionsId := make([]int, len(vtOptions))
	for i := 0; i < len(vtOptions); i++ {
		VoteOptionsId[i] = vtOptions[i].ID
	}
	err2 := s.voiceRepo.DeleteUserVoices(usr.ID, VoteOptionsId)
	if err2 != nil {
		return err2
	}
	var dto = voice.VoiceDTO{
		VoteOptionID: id,
		UserID:       usr.ID,
		Time:         time.Now(),
	}
	err3 := s.voiceRepo.Insert(dto)
	if err3 != nil {
		return err3
	}
	return nil
}

func (s *VoteService) CancelVote(ctx context.Context, id int) error {
	usr := ctx.Value("User").(user.User)
	err := s.voiceRepo.DeleteVoice(usr.ID, id)
	if err != nil {
		return err
	}
	return nil
}
