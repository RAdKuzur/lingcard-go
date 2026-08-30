package userService

import (
	"context"
	"lingcard-go/dictionaries/role"
	"lingcard-go/dictionaries/word"
	"lingcard-go/dto/profile"
	"lingcard-go/dto/request"
	"lingcard-go/models/user"
	"lingcard-go/repositories/courseRepository"
	"lingcard-go/repositories/userRepository"
	"lingcard-go/repositories/wordTranslationRepository"
)

type UserService struct {
	userRepository            *userRepository.UserRepository
	courseRepository          *courseRepository.CourseRepository
	wordTranslationRepository *wordTranslationRepository.WordTranslationRepository
}

func New(userRepository *userRepository.UserRepository, courseRepository *courseRepository.CourseRepository, wordTranslationRepository *wordTranslationRepository.WordTranslationRepository) *UserService {
	return &UserService{
		userRepository:            userRepository,
		courseRepository:          courseRepository,
		wordTranslationRepository: wordTranslationRepository,
	}
}

func (s *UserService) Profile(ctx context.Context) (profile.ProfileDTO, error) {
	usr := ctx.Value("User").(user.User)
	rl := role.RoleDictionary{}.Get(usr.Role)
	courses, err := s.courseRepository.GetCoursesByStatus(usr.ID, []int{word.LEARNING, word.LEARNED})
	if err != nil {
		return profile.ProfileDTO{}, err
	}
	wordTranslationIDs := make([]int, len(courses))
	for i, course := range courses {
		wordTranslationIDs[i] = course.WordTranslationID
	}
	countWordTranslations, err1 := s.wordTranslationRepository.CountNewWords(usr.BaseLanguageID, usr.TargetLanguageID, wordTranslationIDs)
	if err1 != nil {
		return profile.ProfileDTO{}, err1
	}
	learningWords, err2 := s.courseRepository.CountUserStats(usr.ID, word.LEARNING)
	if err2 != nil {
		return profile.ProfileDTO{}, err2
	}
	learnedWords, err3 := s.courseRepository.CountUserStats(usr.ID, word.LEARNED)
	if err3 != nil {
		return profile.ProfileDTO{}, err3
	}
	return profile.ProfileDTO{
		Username:         usr.Name,
		Role:             rl,
		BaseLanguageID:   usr.BaseLanguageID,
		TargetLanguageID: usr.TargetLanguageID,
		NoneWords:        countWordTranslations,
		LearningWords:    learningWords,
		LearnedWords:     learnedWords,
	}, nil
}

func (s *UserService) Update(ctx context.Context, dto request.ProfileUpdateDTO) error {
	usr := ctx.Value("User").(user.User)
	err := s.userRepository.Update(usr.ID, dto)
	if err != nil {
		return err
	}
	return nil
}
