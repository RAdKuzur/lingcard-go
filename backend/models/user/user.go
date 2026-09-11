package user

import (
	"time"
)

type User struct {
	ID               int        `gorm:"column:id;primary_key"`
	Name             string     `gorm:"column:name"`
	Email            *string    `gorm:"column:email;uniqueIndex"`
	EmailVerifiedAt  *time.Time `gorm:"column:email_verified_at"`
	Password         string     `gorm:"column:password"`
	RememberToken    *string    `gorm:"column:remember_token"`
	CreatedAt        time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt        time.Time  `gorm:"column:updated_at;autoUpdateTime"`
	BaseLanguageID   int        `gorm:"column:base_language_id"`
	TargetLanguageID int        `gorm:"column:target_language_id"`
	Role             int        `gorm:"column:role;default:0"`
	IsBanned         bool       `gorm:"column:is_banned;default:false"`
}

func (User) TableName() string {
	return "users"
}
