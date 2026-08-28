package courseService

import (
	"context"
	"lingcard-go/models/user"
	"lingcard-go/repositories/courseRepository"
	"lingcard-go/repositories/wordTranslationRepository"
)

type CourseService struct {
	courseRepository          *courseRepository.CourseRepository
	wordTranslationRepository *wordTranslationRepository.WordTranslationRepository
}

func New(courseRepository *courseRepository.CourseRepository, wordTranslationRepository *wordTranslationRepository.WordTranslationRepository) *CourseService {
	return &CourseService{
		courseRepository:          courseRepository,
		wordTranslationRepository: wordTranslationRepository,
	}
}

func (s *CourseService) ClearProgress(ctx context.Context) {
	User := ctx.Value("User").(user.User)
	s.courseRepository.DeleteCoursesByUserID(User.ID)
}
