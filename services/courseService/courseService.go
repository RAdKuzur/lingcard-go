package courseService

import (
	"context"
	"lingcard-go/dictionaries/level"
	"lingcard-go/dictionaries/word"
	wordProgressDTO "lingcard-go/dto/word"
	"lingcard-go/models/user"
	"lingcard-go/repositories/courseRepository"
	"lingcard-go/repositories/wordRepository"
	"lingcard-go/repositories/wordTranslationRepository"
)

type CourseService struct {
	courseRepository          *courseRepository.CourseRepository
	wordTranslationRepository *wordTranslationRepository.WordTranslationRepository
	wordRepository            *wordRepository.WordRepository
}

func New(courseRepository *courseRepository.CourseRepository, wordTranslationRepository *wordTranslationRepository.WordTranslationRepository, wordRepository *wordRepository.WordRepository) *CourseService {
	return &CourseService{
		courseRepository:          courseRepository,
		wordTranslationRepository: wordTranslationRepository,
		wordRepository:            wordRepository,
	}
}

func (s *CourseService) ClearProgress(ctx context.Context) {
	User := ctx.Value("User").(user.User)
	s.courseRepository.DeleteCoursesByUserID(User.ID)
}

func (s *CourseService) ClearWordProgress(courseID int) {
	s.courseRepository.DeleteWordProgress(courseID)
}

func (s *CourseService) WordsByStatus(ctx context.Context, status, page, limit int, search string) ([]wordProgressDTO.WordProgressDTO, int) {
	var data []wordProgressDTO.WordProgressDTO
	var amountWords int = 0
	User := ctx.Value("User").(user.User)
	switch status {
	case word.NONE:
		{
			courses := s.courseRepository.GetCoursesByStatus(User.ID, []int{word.LEARNING, word.LEARNED})
			wordTranslationIDs := make([]int, len(courses))
			for i, course := range courses {
				wordTranslationIDs[i] = course.WordTranslationID
			}
			wordTranslations := s.wordTranslationRepository.GetSearchNewWords(User.BaseLanguageID, User.TargetLanguageID, page, limit, search, wordTranslationIDs)
			for _, wordTranslation := range wordTranslations {
				Word := s.wordRepository.Find(wordTranslation.WordID)
				Level, _ := level.LevelDictionary{}.Get(Word.Level)
				data = append(data, wordProgressDTO.WordProgressDTO{
					ID:            wordTranslation.ID,
					Text:          Word.Text,
					Translation:   wordTranslation.Translation,
					Transcription: Word.Transcription,
					Level:         Level,
					RepeatTime:    "",
				})
			}
			amountWords = s.wordTranslationRepository.CountSearchNewWords(User.BaseLanguageID, User.TargetLanguageID, search, wordTranslationIDs)
			return data, amountWords
		}
	case word.LEARNING:
		{
			courses := s.courseRepository.GetByStatus(status, User.ID, page, limit, search)
			for _, course := range courses {
				wordTranslation := s.wordTranslationRepository.Find(course.WordTranslationID)
				Word := s.wordRepository.Find(wordTranslation.WordID)
				Level, _ := level.LevelDictionary{}.Get(Word.Level)
				data = append(data, wordProgressDTO.WordProgressDTO{
					ID:            course.ID,
					Text:          Word.Text,
					Translation:   wordTranslation.Translation,
					Transcription: Word.Transcription,
					Level:         Level,
					RepeatTime:    course.LastTimeRepeated.Format("2006-01-02 15:04:05"),
				})
			}
			return data, amountWords
		}
	case word.LEARNED:
		{
			courses := s.courseRepository.GetByStatus(status, User.ID, page, limit, search)

			for _, course := range courses {
				wordTranslation := s.wordTranslationRepository.Find(course.WordTranslationID)
				Word := s.wordRepository.Find(wordTranslation.WordID)
				Level, _ := level.LevelDictionary{}.Get(Word.Level)
				data = append(data, wordProgressDTO.WordProgressDTO{
					ID:            course.ID,
					Text:          Word.Text,
					Translation:   wordTranslation.Translation,
					Transcription: Word.Transcription,
					Level:         Level,
					RepeatTime:    course.LastTimeRepeated.Format("2006-01-02 15:04:05"),
				})
			}
			return data, amountWords
		}
	default:
		amountWords = 0
		return data, amountWords
	}
}
