package course

import "time"

type SimpleCourseDTO struct {
	Status           int       `json:"status"`
	LastTimeRepeated time.Time `json:"last_time_repeated"`
}

type CourseDTO struct {
	Repeat           int       `json:"repeat"`
	Status           int       `json:"status"`
	LastTimeRepeated time.Time `json:"last_time_repeated"`
}

type ExtendedCourseDTO struct {
	WordTranslationID int       `json:"word_translation_id"`
	Repeat            int       `json:"repeat"`
	Status            int       `json:"status"`
	LastTimeRepeated  time.Time `json:"last_time_repeated"`
	UserID            int       `json:"user_id"`
}
