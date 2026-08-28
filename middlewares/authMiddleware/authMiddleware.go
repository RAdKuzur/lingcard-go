package authMiddleware

import (
	"context"
	"gorm.io/gorm"
	"lingcard-go/models/token"
	"lingcard-go/models/user"
	"net/http"
)

type AuthMiddleware struct {
	db *gorm.DB
}

func New(db *gorm.DB) *AuthMiddleware {
	return &AuthMiddleware{db: db}
}

func (m *AuthMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		refreshCookie, err := r.Cookie("refresh_token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		refreshToken := refreshCookie.Value
		var Token token.Token
		err = m.db.Where("refresh_token = ?", refreshToken).First(&Token).Error
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var User user.User
		err = m.db.First(&User, Token.UserID).Error
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "User", User)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
