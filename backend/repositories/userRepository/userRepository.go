package userRepository

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"lingcard-go/dto/request"
	userDTO "lingcard-go/dto/user"
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
func (r *UserRepository) Insert(dto userDTO.UserDTO) error {
	err := r.db.Exec("INSERT INTO users (name, password, email, role, target_language_id, base_language_id, is_banned) VALUES (?, ?, ?, ?, ?, ?, ?)",
		dto.Name, dto.Password, dto.Email, dto.Role, dto.TargetLanguageID, dto.BaseLanguageID, dto.IsBanned).Error
	return err
}

func (r *UserRepository) One(id int) (*user.User, error) {
	var item user.User
	err := r.db.Raw("SELECT * FROM users WHERE id = ?", id).Scan(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *UserRepository) CheckCredentials(name, password string) (bool, error) {
	var item user.User
	err := r.db.Raw("SELECT * FROM users WHERE name = ? AND is_banned = ?", name, false).Scan(&item).Error
	if err != nil {
		return false, err
	}
	if !checkPasswordHash(password, item.Password) {
		return false, nil
	}
	return true, nil
}

func (r *UserRepository) GetUserByCredentials(name, password string) (*user.User, error) {
	var item user.User
	err := r.db.Raw("SELECT * FROM users WHERE name = ? AND is_banned = ?", name, false).Scan(&item).Error

	if err != nil {
		return &user.User{}, err
	}
	if !checkPasswordHash(password, item.Password) {
		return nil, err
	}
	return &item, nil
}

func (r *UserRepository) GetUserByUsername(name string) (*user.User, error) {
	var item user.User
	err := r.db.Raw("SELECT * FROM users WHERE name = ? AND is_banned = ?", name, false).Scan(&item).Error

	if err != nil {
		return nil, err
	}
	return &item, err
}

func (r *UserRepository) Unique(name string) (bool, error) {
	var item user.User
	err := r.db.Raw("SELECT * FROM users WHERE name = ? AND is_banned = ?", name, false).Scan(&item).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}
		return false, err
	}

	return false, nil
}
func (r *UserRepository) Update(id int, dto request.ProfileUpdateDTO) error {
	err := r.db.Exec("UPDATE users SET base_language_id = ?, target_language_id = ? WHERE id = ?", dto.BaseLanguageID, dto.TargetLanguageID, id).Error
	return err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
