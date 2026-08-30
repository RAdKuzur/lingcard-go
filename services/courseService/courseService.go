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

func (s *CourseService) ClearProgress(ctx context.Context) error {
	usr := ctx.Value("User").(user.User)
	err := s.courseRepository.DeleteCoursesByUserID(usr.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *CourseService) ClearWordProgress(courseID int) error {
	err := s.courseRepository.DeleteWordProgress(courseID)
	if err != nil {
		return err
	}
	return nil
}

func (s *CourseService) WordsByStatus(ctx context.Context, status, page, limit int, search string) ([]wordProgressDTO.WordProgressDTO, int, error) {
	var data = make([]wordProgressDTO.WordProgressDTO, 0)
	var amountWords = 0
	usr := ctx.Value("User").(user.User)
	switch status {
	case word.NONE:
		{
			courses, err := s.courseRepository.GetCoursesByStatus(usr.ID, []int{word.LEARNING, word.LEARNED})
			if err != nil {
				return nil, 0, err
			}
			wordTranslationIDs := make([]int, len(courses))
			for i, course := range courses {
				wordTranslationIDs[i] = course.WordTranslationID
			}
			wordTranslations, err2 := s.wordTranslationRepository.GetSearchNewWords(usr.BaseLanguageID, usr.TargetLanguageID, page, limit, search, wordTranslationIDs)
			if err2 != nil {
				return nil, 0, err2
			}
			for _, wordTranslation := range wordTranslations {
				wrd, err3 := s.wordRepository.Find(wordTranslation.WordID)
				if err3 != nil {
					return nil, 0, err3
				}
				lvl := level.LevelDictionary{}.Get(wrd.Level)
				data = append(data, wordProgressDTO.WordProgressDTO{
					ID:            wordTranslation.ID,
					Text:          wrd.Text,
					Translation:   wordTranslation.Translation,
					Transcription: wrd.Transcription,
					Level:         lvl,
					RepeatTime:    "",
				})
			}
			amountWords, err = s.wordTranslationRepository.CountSearchNewWords(usr.BaseLanguageID, usr.TargetLanguageID, search, wordTranslationIDs)
			if err != nil {
				return nil, 0, err
			}
			return data, amountWords, nil
		}
	case word.LEARNING:
		{
			courses, err4 := s.courseRepository.GetByStatus(status, usr.ID, page, limit, search)
			if err4 != nil {
				return nil, 0, err4
			}
			for _, course := range courses {
				wordTranslation, err5 := s.wordTranslationRepository.Find(course.WordTranslationID)
				if err5 != nil {
					return nil, 0, err5
				}
				wrd, err6 := s.wordRepository.Find(wordTranslation.WordID)
				if err6 != nil {
					return nil, 0, err6
				}
				lvl := level.LevelDictionary{}.Get(wrd.Level)
				data = append(data, wordProgressDTO.WordProgressDTO{
					ID:            course.ID,
					Text:          wrd.Text,
					Translation:   wordTranslation.Translation,
					Transcription: wrd.Transcription,
					Level:         lvl,
					RepeatTime:    course.LastTimeRepeated.Format("2006-01-02 15:04:05"),
				})
			}
			amountWords, _ = s.courseRepository.CountByStatus(word.LEARNING, usr.ID, search)
			return data, amountWords, nil
		}
	case word.LEARNED:
		{
			courses, err7 := s.courseRepository.GetByStatus(status, usr.ID, page, limit, search)
			if err7 != nil {
				return nil, 0, err7
			}
			for _, course := range courses {
				wordTranslation, err8 := s.wordTranslationRepository.Find(course.WordTranslationID)
				if err8 != nil {
					return nil, 0, err8
				}
				wrd, err9 := s.wordRepository.Find(wordTranslation.WordID)
				if err9 != nil {
					return nil, 0, err9
				}
				lvl := level.LevelDictionary{}.Get(wrd.Level)
				data = append(data, wordProgressDTO.WordProgressDTO{
					ID:            course.ID,
					Text:          wrd.Text,
					Translation:   wordTranslation.Translation,
					Transcription: wrd.Transcription,
					Level:         lvl,
					RepeatTime:    course.LastTimeRepeated.Format("2006-01-02 15:04:05"),
				})
			}
			amountWords, _ = s.courseRepository.CountByStatus(word.LEARNED, usr.ID, search)
			return data, amountWords, nil
		}
	default:
		amountWords = 0
		return data, amountWords, nil
	}
}

func (s *CourseService) NewWord(ctx context.Context) (wordProgressDTO.WordTrainingDTO, error) {
	usr := ctx.Value("User").(user.User)
	var dto wordProgressDTO.WordTrainingDTO
	count, err := s.courseRepository.CountRepeatedWords(usr.ID)
	if err != nil {
		return wordProgressDTO.WordTrainingDTO{}, err
	}

	if count > 0 {
		course, err2 := s.courseRepository.GetOldLearningWords(usr.ID)
		if err2 != nil {
			return wordProgressDTO.WordTrainingDTO{}, err2
		}
		wrdTranslation, err3 := s.wordTranslationRepository.Find(course.WordTranslationID)
		if err3 != nil {
			return wordProgressDTO.WordTrainingDTO{}, err3
		}
		wrd, err4 := s.wordRepository.Find(wrdTranslation.WordID)
		if err4 != nil {
			return wordProgressDTO.WordTrainingDTO{}, err4
		}
		lvl := level.LevelDictionary{}.Get(wrd.Level)
		dto = wordProgressDTO.WordTrainingDTO{
			ID:            course.WordTranslationID,
			Text:          wrd.Text,
			Translation:   wrdTranslation.Translation,
			Transcription: wrd.Transcription,
			Level:         lvl,
			Status:        course.Status,
			Repeat:        course.Repeat,
		}
	} else {
		courses, err5 := s.courseRepository.GetCoursesByStatus(usr.ID, []int{word.LEARNING, word.LEARNED})
		if err5 != nil {
			return wordProgressDTO.WordTrainingDTO{}, err5
		}
		wordTranslationIDs := make([]int, len(courses))
		for i, course := range courses {
			wordTranslationIDs[i] = course.WordTranslationID
		}
		wrdTranslation, err6 := s.courseRepository.GetNewWord(usr.BaseLanguageID, usr.TargetLanguageID, wordTranslationIDs)
		if err6 != nil {
			return wordProgressDTO.WordTrainingDTO{}, err6
		}
		if wrdTranslation.ID != 0 {
			wrd, err7 := s.wordRepository.Find(wrdTranslation.WordID)
			if err7 != nil {
				return wordProgressDTO.WordTrainingDTO{}, err7
			}
			lvl := level.LevelDictionary{}.Get(wrd.Level)
			dto = wordProgressDTO.WordTrainingDTO{
				ID:            wrdTranslation.ID,
				Text:          wrd.Text,
				Translation:   wrdTranslation.Translation,
				Transcription: wrd.Transcription,
				Level:         lvl,
				Status:        word.NONE,
				Repeat:        0,
			}
		} else {
			dto = wordProgressDTO.WordTrainingDTO{}
		}
	}
	return dto, nil
}

func (s *CourseService) RepeatWord(ctx context.Context, id int, status bool) error {
	User := ctx.Value("User").(user.User)
	course, err := s.courseRepository.GetCourseByWordTranslationIDAndUserID(id, User.ID)
	if err != nil {
		return err
	}
	if course.ID != 0 {
		if status == true {
			dto := courseDTO.CourseDTO{
				Repeat:           course.Repeat + 1,
				Status:           getStatus(course.Repeat),
				LastTimeRepeated: Repeat(course.Repeat),
			}
			err2 := s.courseRepository.Update(course.ID, dto)
			if err2 != nil {
				return err2
			}
		} else {
			dto := courseDTO.SimpleCourseDTO{
				Status:           word.LEARNING,
				LastTimeRepeated: time.Now().Add(10 * time.Minute),
			}
			err3 := s.courseRepository.SimpleUpdate(course.ID, dto)
			if err3 != nil {
				return err3
			}
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
			err4 := s.courseRepository.ExtendedInsert(dto)
			if err4 != nil {
				return err4
			}
		} else {
			dto := courseDTO.ExtendedCourseDTO{
				WordTranslationID: id,
				UserID:            User.ID,
				Repeat:            0,
				Status:            word.LEARNING,
				LastTimeRepeated:  Repeat(0),
			}
			err5 := s.courseRepository.ExtendedInsert(dto)
			if err5 != nil {
				return err5
			}
		}
	}
	return nil
}

func (s *CourseService) Teachable(ctx context.Context) (string, int, error) {
	usr := ctx.Value("User").(user.User)
	language, err := s.languageRepository.Find(usr.TargetLanguageID)
	if err != nil {
		return "", 0, err
	}
	allWordAmount, err2 := s.wordTranslationRepository.CountByTargetLanguageIdAndBaseLanguageId(usr.BaseLanguageID, usr.TargetLanguageID, "")
	if err2 != nil {
		return "", 0, err2
	}
	learnedWords, err3 := s.courseRepository.GetCoursesByStatus(usr.ID, []int{word.LEARNED})
	if err3 != nil {
		return "", 0, err3
	}
	amountLearningWords, err4 := s.courseRepository.CountRepeatedWords(usr.ID)
	if err4 != nil {
		return "", 0, err4
	}
	repeatedWords, err5 := s.courseRepository.GetCoursesByStatus(usr.ID, []int{word.LEARNED, word.LEARNING})
	if err5 != nil {
		return "", 0, err5
	}
	if len(learnedWords) == allWordAmount {
		return language.Code, word.LEARNED, nil
	} else {
		if amountLearningWords == 0 && allWordAmount == len(repeatedWords) {
			return language.Code, word.LEARNING, nil
		} else {
			return language.Code, word.NONE, nil
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
