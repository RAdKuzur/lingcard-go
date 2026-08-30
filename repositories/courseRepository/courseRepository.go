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
func (r *CourseRepository) GetByStatus(status, userID, page, limit int, search string) ([]course.Course, error) {
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
		return nil, err
	}
	return courses, err
}

func (r *CourseRepository) GetCoursesByStatus(userID int, statuses []int) ([]course.Course, error) {
	var courses []course.Course
	err := r.db.Raw("SELECT * FROM courses WHERE user_id = ? AND status IN (?)", userID, statuses).Scan(&courses).Error
	return courses, err
}

func (r *CourseRepository) CountByStatus(status, userID int, search string) (int, error) {
	var count int
	query := `
        SELECT COUNT(*)
        FROM courses
        JOIN word_translations ON courses.word_translation_id = word_translations.id
        JOIN words ON word_translations.word_id = words.id
        WHERE words.text LIKE ?
        AND courses.status = ?
        AND courses.user_id = ?
    `
	err := r.db.Raw(query, "%"+search+"%", status, userID).Scan(&count).Error
	if err != nil {
		return 0, err
	}
	return count, err
}

func (r *CourseRepository) CountUserStats(userID int, status int) (int, error) {
	var count int
	err := r.db.Raw("SELECT COUNT(*) FROM courses WHERE user_id = ? AND status = ?", userID, status).Scan(&count).Error
	return count, err
}

func (r *CourseRepository) DeleteCoursesByUserID(userID int) error {
	err := r.db.Exec("DELETE FROM courses WHERE user_id = ?", userID).Error
	return err
}

func (r *CourseRepository) DeleteWordProgress(id int) error {
	err := r.db.Exec("DELETE FROM courses WHERE id = ?", id).Error
	return err
}

func (r *CourseRepository) GetCourseByWordTranslationIDAndUserID(ID int, UserID int) (course.Course, error) {
	var item course.Course
	err := r.db.Raw("SELECT * FROM courses WHERE word_translation_id = ? AND user_id = ? LIMIT 1", ID, UserID).Scan(&item).Error
	return item, err
}

func (r *CourseRepository) SimpleUpdate(id int, dto courseDTO.SimpleCourseDTO) error {
	err := r.db.Exec("UPDATE courses SET status = ?, last_time_repeated = ? WHERE id = ?", dto.Status, dto.LastTimeRepeated, id).Error
	return err
}

func (r *CourseRepository) Update(id int, dto courseDTO.CourseDTO) error {
	err := r.db.Exec("UPDATE courses SET repeat = ?, status = ?, last_time_repeated = ? WHERE id = ?", dto.Repeat, dto.Status, dto.LastTimeRepeated, id).Error
	return err
}

func (r *CourseRepository) ExtendedInsert(dto courseDTO.ExtendedCourseDTO) error {
	err := r.db.Exec("INSERT INTO courses (word_translation_id, user_id, repeat, status, last_time_repeated) VALUES (?, ?, ?, ?, ?)",
		dto.WordTranslationID, dto.UserID, dto.Repeat, dto.Status, dto.LastTimeRepeated).Error
	return err
}

func (r *CourseRepository) CountRepeatedWords(userID int) (int, error) {
	var count int
	err := r.db.Raw("SELECT COUNT(*) FROM courses WHERE user_id = ? AND status = ? AND last_time_repeated < ?", userID, word.LEARNING, time.Now()).Scan(&count).Error
	return count, err
}

func (r *CourseRepository) GetOldLearningWords(userID int) (course.Course, error) {
	var item course.Course
	err := r.db.
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
		First(&item).Error
	return item, err
}
func (r *CourseRepository) GetNewWord(baseLangID, targetLangID int, exceptID []int) (*wordModel.WordTranslation, error) {
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

	err := query.Order("RANDOM()").First(&wordTranslation).Error

	return &wordTranslation, err
}
