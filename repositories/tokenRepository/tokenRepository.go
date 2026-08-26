package tokenRepository

import (
	"gorm.io/gorm"
	"lingcard-go/models/token"
	"os"
	"strconv"
	"time"
)

type TokenRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

func (r *TokenRepository) CreateToken(refreshToken string, userId int, ip string, userAgent string) error {
	var minutes, _ = strconv.Atoi(os.Getenv("REFRESH_TOKEN_TIME_EXPIRE"))
	var item = token.Token{
		RefreshToken: refreshToken,
		UserAgent:    userAgent,
		UserID:       userId,
		IPAddress:    ip,
		IsRevoked:    false,
		ExpiresAt:    time.Now().Add(time.Minute * time.Duration(minutes)).String(),
	}
	err := r.db.Create(&item).Error
	return err
}

func (r *TokenRepository) GetByRefreshToken(refreshToken string) (token.Token, error) {
	var item token.Token
	err := r.db.Where("refresh_token = ?", refreshToken).First(&item).Error
	if err != nil {
		return item, err
	}
	return item, nil
}

func (r *TokenRepository) Delete(id int) error {
	err := r.db.Where("id = ?", id).Delete(&token.Token{}).Error
	return err
}
