package courseService

import (
	"context"
	"lingcard-go/dictionaries/word"
	wordProgressDTO "lingcard-go/dto/word"
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

func (s *CourseService) ClearWordProgress(courseID int) {
	s.courseRepository.DeleteWordProgress(courseID)
}

func (s *CourseService) WordsByStatus(ctx context.Context, status, page, limit int, search string) []wordProgressDTO.WordProgressDTO {
	var data []wordProgressDTO.WordProgressDTO
	User := ctx.Value("User").(user.User)
	switch status {
	case word.NONE:
		{

		}
	case word.LEARNING:
		{
			courses := s.courseRepository.GetByStatus(status, User.ID, page, limit, search)
			for _, course := range courses {
				data = append(data, wordProgressDTO.WordProgressDTO{
					ID: course.ID,
				})
			}
		}

	}
	return data
}
