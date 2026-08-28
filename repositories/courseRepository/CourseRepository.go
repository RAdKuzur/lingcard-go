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
