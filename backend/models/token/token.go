package token

type Token struct {
	ID           int    `gorm:"column:id;primary_key"`
	RefreshToken string `gorm:"column:refresh_token;uniqueIndex;not null"`
	ExpiresAt    string `gorm:"column:expires_at;not null"`
	UserID       int    `gorm:"column:user_id;index;not null"`
	IPAddress    string `gorm:"column:ip_address"`
	UserAgent    string `gorm:"column:user_agent"`
	IsRevoked    bool   `gorm:"column:is_revoked;default:false;index"`
}

func (Token) TableName() string {
	return "tokens"
}
