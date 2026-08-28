package courseRepository

import (
	"gorm.io/gorm"
	"lingcard-go/models/course"
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
