package courseService

import (
	"context"
	"lingcard-go/dictionaries/level"
	"lingcard-go/dictionaries/word"
	courseDTO "lingcard-go/dto/course"
	wordProgressDTO "lingcard-go/dto/word"
	"lingcard-go/models/user"
	"lingcard-go/repositories/courseRepository"
	"lingcard-go/repositories/languageRepository"
	"lingcard-go/repositories/wordRepository"
	"lingcard-go/repositories/wordTranslationRepository"
	"time"
)

const REPEAT_TIME int = 6

type CourseService struct {
	courseRepository          *courseRepository.CourseRepository
	wordTranslationRepository *wordTranslationRepository.WordTranslationRepository
	wordRepository            *wordRepository.WordRepository
	languageRepository        *languageRepository.LanguageRepository
}

func New(courseRepository *courseRepository.CourseRepository, wordTranslationRepository *wordTranslationRepository.WordTranslationRepository, wordRepository *wordRepository.WordRepository, languageRepository *languageRepository.LanguageRepository) *CourseService {
	return &CourseService{
		courseRepository:          courseRepository,
		wordTranslationRepository: wordTranslationRepository,
		wordRepository:            wordRepository,
		languageRepository:        languageRepository,
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

func (s *CourseService) NewWord(ctx context.Context) wordProgressDTO.WordTrainingDTO {
	User := ctx.Value("User").(user.User)
	count := s.courseRepository.CountRepeatedWords(User.ID)
	var dto wordProgressDTO.WordTrainingDTO
	if count > 0 {
		course := s.courseRepository.GetOldLearningWords(User.ID)
		WordTranslation := s.wordTranslationRepository.Find(course.WordTranslationID)
		Word := s.wordRepository.Find(WordTranslation.WordID)
		Level, _ := level.LevelDictionary{}.Get(Word.Level)
		dto = wordProgressDTO.WordTrainingDTO{
			ID:            course.WordTranslationID,
			Text:          Word.Text,
			Translation:   WordTranslation.Translation,
			Transcription: Word.Transcription,
			Level:         Level,
			Status:        course.Status,
			Repeat:        course.Repeat,
		}
	} else {
		courses := s.courseRepository.GetCoursesByStatus(User.ID, []int{word.LEARNING, word.LEARNED})
		wordTranslationIDs := make([]int, len(courses))
		for i, course := range courses {
			wordTranslationIDs[i] = course.WordTranslationID
		}
		WordTranslation := s.courseRepository.GetNewWord(User.BaseLanguageID, User.TargetLanguageID, wordTranslationIDs)
		if WordTranslation.ID != 0 {
			Word := s.wordRepository.Find(WordTranslation.WordID)
			Level, _ := level.LevelDictionary{}.Get(Word.Level)
			dto = wordProgressDTO.WordTrainingDTO{
				ID:            WordTranslation.ID,
				Text:          Word.Text,
				Translation:   WordTranslation.Translation,
				Transcription: Word.Transcription,
				Level:         Level,
				Status:        word.NONE,
				Repeat:        0,
			}
		} else {
			dto = wordProgressDTO.WordTrainingDTO{}
		}
	}
	return dto
}

func (s *CourseService) RepeatWord(ctx context.Context, id int, status bool) {
	User := ctx.Value("User").(user.User)
	course := s.courseRepository.GetCourseByWordTranslationIDAndUserID(id, User.ID)
	if course.ID != 0 {
		if status == true {
			dto := courseDTO.CourseDTO{
				Repeat:           course.Repeat + 1,
				Status:           getStatus(course.Repeat),
				LastTimeRepeated: Repeat(course.Repeat),
			}
			s.courseRepository.Update(course.ID, dto)
		} else {
			dto := courseDTO.SimpleCourseDTO{
				Status:           word.LEARNING,
				LastTimeRepeated: time.Now().Add(10 * time.Minute),
			}
			s.courseRepository.SimpleUpdate(course.ID, dto)
		}
	} else {
		if status == true {
			dto := courseDTO.ExtendedCourseDTO{
				WordTranslationID: id,
				UserID:            User.ID,
				Repeat:            0,
				Status:            word.LEARNED,
				LastTimeRepeated:  Repeat(0),
			}
			s.courseRepository.ExtendedUpdate(course.ID, dto)
		} else {
			dto := courseDTO.ExtendedCourseDTO{
				WordTranslationID: id,
				UserID:            User.ID,
				Repeat:            0,
				Status:            word.LEARNING,
				LastTimeRepeated:  Repeat(0),
			}
			s.courseRepository.ExtendedUpdate(course.ID, dto)
		}
	}
}

func (s *CourseService) Teachable(ctx context.Context) (string, int) {
	User := ctx.Value("User").(user.User)
	language := s.languageRepository.Find(User.TargetLanguageID)
	allWordAmount := s.wordTranslationRepository.CountByTargetLanguageIdAndBaseLanguageId(User.BaseLanguageID, User.TargetLanguageID, "")
	amountLearnedWords := len(s.courseRepository.GetCoursesByStatus(User.ID, []int{word.LEARNED}))
	amountLearningWords := s.courseRepository.CountRepeatedWords(User.ID)
	if amountLearnedWords == allWordAmount {
		return language.Code, word.LEARNED
	} else {
		if amountLearningWords == 0 && allWordAmount == len(s.courseRepository.GetCoursesByStatus(User.ID, []int{word.LEARNED, word.LEARNING})) {
			return language.Code, word.LEARNING
		} else {
			return language.Code, word.NONE
		}
	}
}

func Repeat(count int) time.Time {
	var duration time.Duration
	switch count {
	case 0:
		duration = 30 * time.Minute
	case 1:
		duration = 4 * time.Hour
	case 2:
		duration = 24 * time.Hour
	case 3:
		duration = 3 * 24 * time.Hour
	case 4:
		duration = 7 * 24 * time.Hour
	case 5:
		duration = 30 * 24 * time.Hour
	case 6:
		duration = 60 * 24 * time.Hour
	default:
		duration = 90 * 24 * time.Hour
	}
	return time.Now().Add(duration)
}

func getStatus(repeat int) int {
	if repeat > REPEAT_TIME {
		return word.LEARNED
	}
	return word.LEARNING
}
