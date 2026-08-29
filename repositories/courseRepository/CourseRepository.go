package courseRepository

import (
	"gorm.io/gorm"
	"lingcard-go/dictionaries/word"
	courseDTO "lingcard-go/dto/course"
	"lingcard-go/models/course"
	wordModel "lingcard-go/models/word"
	"time"
)

type CourseRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}
func (r *CourseRepository) GetByStatus(status, userID, page, limit int, search string) []course.Course {
	var courses []course.Course
	query := r.db.Table("courses").
		Select("courses.*").
		Joins("JOIN word_translations ON courses.word_translation_id = word_translations.id").
		Joins("JOIN words ON word_translations.word_id = words.id").
		Where("courses.user_id = ?", userID).
		Where("courses.status = ?", status)
	if search != "" {
		query = query.Where("words.text LIKE ?", "%"+search+"%")
	}
	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Find(&courses).Error
	if err != nil {
		return nil
	}
	return courses
}

func (r *CourseRepository) GetCoursesByStatus(userID int, statuses []int) []course.Course {
	var courses []course.Course
	r.db.Raw("SELECT * FROM courses WHERE user_id = ? AND status IN (?)", userID, statuses).Scan(&courses)
	return courses
}

func (r *CourseRepository) CountUserStats(userID int, status int) int {
	var count int
	r.db.Raw("SELECT COUNT(*) FROM courses WHERE user_id = ? AND status = ?", userID, status).Scan(&count)
	return count
}

func (r *CourseRepository) DeleteCoursesByUserID(userID int) {
	r.db.Exec("DELETE FROM courses WHERE user_id = ?", userID)
}

func (r *CourseRepository) DeleteWordProgress(id int) {
	r.db.Exec("DELETE FROM courses WHERE id = ?", id)
}

func (r *CourseRepository) GetCourseByWordTranslationIDAndUserID(ID int, UserID int) course.Course {
	var item course.Course
	r.db.Raw("SELECT * FROM courses WHERE word_translation_id = ? AND user_id = ? LIMIT 1", ID, UserID).Scan(&item)
	return item
}

func (r *CourseRepository) SimpleUpdate(id int, dto courseDTO.SimpleCourseDTO) {
	r.db.Exec("UPDATE courses SET (status, last_time_repeated) VALUES (?, ?) WHERE id = ?", dto.Status, dto.LastTimeRepeated, id)
}

func (r *CourseRepository) Update(id int, dto courseDTO.CourseDTO) {
	r.db.Exec("UPDATE courses SET (repeat, status, last_time_repeated) VALUES (?, ?, ?) WHERE id = ?", dto.Repeat, dto.Status, dto.LastTimeRepeated, id)
}

func (r *CourseRepository) ExtendedUpdate(id int, dto courseDTO.ExtendedCourseDTO) {
	r.db.Exec("UPDATE courses SET (word_translation_id, user_id, repeat, status, last_time_repeated) VALUES (?, ?, ?, ?, ?)  WHERE id = ?",
		dto.WordTranslationID, dto.UserID, dto.Repeat, dto.Status, dto.LastTimeRepeated, id)
}

func (r *CourseRepository) CountRepeatedWords(userID int) int {
	var count int
	r.db.Exec("SELECT COUNT(*) FROM courses WHERE user_id = ? AND status = ? AND last_time_repeated < ?", userID, word.LEARNING, time.Now()).Scan(&count)
	return count
}

func (r *CourseRepository) GetOldLearningWords(userID int) course.Course {
	var item course.Course
	r.db.
		Table("courses").
		Select("courses.*").
		Joins("JOIN word_translations ON courses.word_translation_id = word_translations.id").
		Joins("JOIN words ON word_translations.word_id = words.id").
		Where("courses.user_id = ?", userID).
		Where("courses.last_time_repeated < ?", time.Now()).
		Where("courses.status = ?", word.LEARNING).
		Order("courses.status DESC").
		Order("words.level ASC").
		Order("courses.last_time_repeated DESC").
		First(&item)
	return item
}
func (r *CourseRepository) GetNewWord(baseLangID, targetLangID int, exceptID []int) *wordModel.WordTranslation {
	var wordTranslation wordModel.WordTranslation
	query := r.db.
		Table("word_translations").
		Select("word_translations.*").
		Joins("JOIN words ON words.id = word_translations.word_id").
		Where("word_translations.target_language_id = ?", baseLangID).
		Where("words.language_id = ?", targetLangID)
	if len(exceptID) > 0 {
		query = query.Where("words.id NOT IN ?", exceptID)
	}

	query.Order("RANDOM()").First(&wordTranslation)

	return &wordTranslation
}
