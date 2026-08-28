package userRepository

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"lingcard-go/dto/request"
	"lingcard-go/models/user"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
func (r *UserRepository) Insert(User user.User) error {
	err := r.db.Create(&User).Error
	return err
}

func (r *UserRepository) One(id int) (*user.User, error) {
	var item user.User
	err := r.db.First(&item, id).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *UserRepository) CheckCredentials(name, password string) (bool, error) {
	var User user.User
	err := r.db.Where("name = ? AND is_banned = ?", name, false).
		First(&User).Error

	if err != nil {
		return false, err
	}
	if !checkPasswordHash(password, User.Password) {
		return false, nil
	}
	return true, nil
}

func (r *UserRepository) GetUserByCredentials(name, password string) (*user.User, error) {
	var User user.User
	err := r.db.Where("name = ? AND is_banned = ?", name, false).
		First(&User).Error

	if err != nil {
		return &user.User{}, err
	}
	if !checkPasswordHash(password, User.Password) {
		return &user.User{}, err
	}
	return &User, nil
}

func (r *UserRepository) GetUserByUsername(name string) (*user.User, error) {
	var User user.User
	err := r.db.Where("name = ? AND is_banned = ?", name, false).
		First(&User).Error

	if err != nil {
		return &user.User{}, err
	}
	return &User, err
}

func (r *UserRepository) Unique(name string) (bool, error) {
	var User user.User
	err := r.db.Where("name = ? AND is_banned = ?", name, false).
		First(&User).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, err
	}

	return false, nil
}
func (r *UserRepository) Update(id int, dto request.ProfileUpdateDTO) {
	r.db.Exec("UPDATE users SET base_language_id = ?, target_language_id = ? WHERE id = ?", dto.BaseLanguageID, dto.TargetLanguageID, id)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
