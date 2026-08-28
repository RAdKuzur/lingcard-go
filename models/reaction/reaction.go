package reaction

type Reaction struct {
	ID     int `gorm:"column:id"`
	UserID int `gorm:"column:user_id"`
	PostID int `gorm:"column:post_id"`
	Status int `gorm:"column:status"`
}
