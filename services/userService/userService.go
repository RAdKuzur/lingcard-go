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

func (s *UserService) Profile(ctx context.Context) profile.ProfileDTO {
	User := ctx.Value("User").(user.User)
	Role, _ := role.RoleDictionary{}.Get(User.Role)
	courses := s.courseRepository.GetCoursesByStatus(User.ID, []int{word.LEARNING, word.LEARNED})

	wordTranslationIDs := make([]int, len(courses))
	for i, course := range courses {
		wordTranslationIDs[i] = course.WordTranslationID
	}
	countWordTranslations := s.wordTranslationRepository.CountNewWords(User.BaseLanguageID, User.TargetLanguageID, wordTranslationIDs)
	learningWords := s.courseRepository.CountUserStats(User.ID, word.LEARNING)
	learnedWords := s.courseRepository.CountUserStats(User.ID, word.LEARNED)
	return profile.ProfileDTO{
		Username:         User.Name,
		Role:             Role,
		BaseLanguageId:   User.BaseLanguageID,
		TargetLanguageId: User.TargetLanguageID,
		NoneWords:        countWordTranslations,
		LearningWords:    learningWords,
		LearnedWords:     learnedWords,
	}
}

func (s *UserService) Update(ctx context.Context, dto request.ProfileUpdateDTO) {
	User := ctx.Value("User").(user.User)
	s.userRepository.Update(User.ID, dto)
}
