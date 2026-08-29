package authService

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"lingcard-go/dictionaries/role"
	"lingcard-go/dto/request"
	"lingcard-go/dto/token"
	"lingcard-go/helpers/auth"
	"lingcard-go/models/user"
	"lingcard-go/repositories/tokenRepository"
	"lingcard-go/repositories/userRepository"
	"net/http"
)

type AuthService struct {
	userRepository  *userRepository.UserRepository
	tokenRepository *tokenRepository.TokenRepository
}

func New(userRepository *userRepository.UserRepository, tokenRepository *tokenRepository.TokenRepository) *AuthService {
	return &AuthService{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
	}
}

func (s *AuthService) Login(name, password string, r *http.Request) (token.TokenDTO, error) {
	ipAddress := r.Header.Get("X-Real-Ip")
	userAgent := r.Header.Get("User-Agent")
	success, err := s.userRepository.CheckCredentials(name, password)
	if err != nil {
		return token.TokenDTO{
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}
	if !success {
		return token.TokenDTO{
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}
	User, err := s.userRepository.GetUserByCredentials(name, password)
	if err != nil {
		return token.TokenDTO{
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}
	accessToken, _ := auth.GenerateAccessToken(User.Name)
	refreshToken, err := auth.GenerateRefreshToken(User.Name)
	if err == nil {
		_ = s.tokenRepository.CreateToken(refreshToken, User.ID, ipAddress, userAgent)
	}
	return token.TokenDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, err
}

func (s *AuthService) Logout(refreshToken string) error {
	tokenRow, err := s.tokenRepository.GetByRefreshToken(refreshToken)
	if err != nil {
		return err
	}
	if tokenRow.ID != 0 {
		err = s.tokenRepository.Delete(tokenRow.ID)
	}
	return err
}

func (s *AuthService) Refresh(refreshToken string, r *http.Request) (token.TokenDTO, error) {
	ipAddress := r.Header.Get("X-Real-Ip")
	userAgent := r.Header.Get("User-Agent")
	tokenRow, err := s.tokenRepository.GetByRefreshToken(refreshToken)
	if err != nil {
		return token.TokenDTO{
			AccessToken:  "",
			RefreshToken: "",
		}, err
	}

	User, err := s.userRepository.One(tokenRow.UserID)

	if User != nil && User.IsBanned == false {
		if err != nil {
			return token.TokenDTO{
				AccessToken:  "",
				RefreshToken: "",
			}, err
		}

		newAccessToken, _ := auth.GenerateAccessToken(User.Name)
		newRefreshToken, err := auth.GenerateRefreshToken(User.Name)

		if err == nil {
			_ = s.tokenRepository.CreateToken(newRefreshToken, User.ID, ipAddress, userAgent)
		}
		_ = s.tokenRepository.Delete(tokenRow.ID)
		return token.TokenDTO{
			AccessToken:  newAccessToken,
			RefreshToken: newRefreshToken,
		}, err
	}

	return token.TokenDTO{
		AccessToken:  "",
		RefreshToken: "",
	}, nil
}

func (s *AuthService) Register(dto request.RegisterDTO, r *http.Request) (token.TokenDTO, error) {
	unique, _ := s.userRepository.Unique(dto.Name)
	if unique {
		ipAddress := r.Header.Get("X-Real-Ip")
		userAgent := r.Header.Get("User-Agent")
		password, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		if err != nil {
			return token.TokenDTO{}, err
		}
		err = s.userRepository.Insert(user.User{
			Name:             dto.Name,
			Password:         string(password),
			Email:            nil,
			Role:             role.RoleUser,
			TargetLanguageID: dto.TargetLang,
			BaseLanguageID:   dto.BaseLangId,
			IsBanned:         false,
		})
		if err != nil {
			return token.TokenDTO{
				AccessToken:  "",
				RefreshToken: "",
			}, err
		}

		User, err := s.userRepository.GetUserByCredentials(dto.Name, dto.Password)
		if err != nil {
			return token.TokenDTO{
				AccessToken:  "",
				RefreshToken: "",
			}, err
		}
		accessToken, _ := auth.GenerateAccessToken(User.Name)
		refreshToken, err := auth.GenerateRefreshToken(User.Name)
		if err == nil {
			_ = s.tokenRepository.CreateToken(refreshToken, User.ID, ipAddress, userAgent)
		}
		return token.TokenDTO{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, err
	}
	return token.TokenDTO{
		AccessToken:  "",
		RefreshToken: "",
	}, errors.New("not unique username")
}

func (s *AuthService) User(accessToken string) (request.AuthUserDTO, error) {
	claims, err := auth.ParseToken(accessToken)
	if err != nil {
		return request.AuthUserDTO{
			Username: "",
			Role:     "",
		}, err
	}
	username := claims["username"].(string)

	User, err := s.userRepository.GetUserByUsername(username)
	if err != nil {
		return request.AuthUserDTO{
			Username: "",
			Role:     "",
		}, err
	}
	role, _ := role.RoleDictionary{}.Get(User.Role)

	return request.AuthUserDTO{
		Username: username,
		Role:     role,
	}, err
}

func (s *AuthService) GetAuthUser(name, password string) request.AuthUserDTO {
	User, _ := s.userRepository.GetUserByCredentials(name, password)
	Role, _ := role.RoleDictionary{}.Get(User.Role)
	var dto = request.AuthUserDTO{
		Username: User.Name,
		Role:     Role,
	}
	return dto
}
