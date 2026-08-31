package authService

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"lingcard-go/dictionaries/role"
	"lingcard-go/dto/request"
	"lingcard-go/dto/token"
	userDTO "lingcard-go/dto/user"
	"lingcard-go/helpers/auth"
	"lingcard-go/repositories/tokenRepository"
	"lingcard-go/repositories/userRepository"
	"log"
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
		}, errors.New("not valid credentials")
	}
	User, err2 := s.userRepository.GetUserByCredentials(name, password)
	if err2 != nil {
		return token.TokenDTO{
			AccessToken:  "",
			RefreshToken: "",
		}, err2
	}
	accessToken, _ := auth.GenerateAccessToken(User.Name)
	refreshToken, err3 := auth.GenerateRefreshToken(User.Name)
	if err3 == nil {
		log.Print(refreshToken, "----", User.ID, "----", ipAddress, "----", userAgent)
		err4 := s.tokenRepository.CreateToken(refreshToken, User.ID, ipAddress, userAgent)
		if err4 != nil {
			return token.TokenDTO{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			}, err4
		}
	}
	return token.TokenDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, err3
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

	User, err2 := s.userRepository.One(tokenRow.UserID)

	if User != nil && User.IsBanned == false {
		if err2 != nil {
			return token.TokenDTO{
				AccessToken:  "",
				RefreshToken: "",
			}, err2
		}

		newAccessToken, _ := auth.GenerateAccessToken(User.Name)
		newRefreshToken, err3 := auth.GenerateRefreshToken(User.Name)

		if err3 == nil {
			errToken := s.tokenRepository.CreateToken(newRefreshToken, User.ID, ipAddress, userAgent)
			if errToken != nil {
				return token.TokenDTO{
					AccessToken:  "",
					RefreshToken: "",
				}, errToken
			}
		}
		err4 := s.tokenRepository.Delete(tokenRow.ID)
		if err4 != nil {
			return token.TokenDTO{
				AccessToken:  "",
				RefreshToken: "",
			}, err4
		}
		return token.TokenDTO{
			AccessToken:  newAccessToken,
			RefreshToken: newRefreshToken,
		}, nil
	}

	return token.TokenDTO{
		AccessToken:  "",
		RefreshToken: "",
	}, errors.New("user is banned")
}

func (s *AuthService) Register(dto request.RegisterDTO, r *http.Request) (token.TokenDTO, error) {
	unique, _ := s.userRepository.Unique(dto.Name)
	if !unique {
		ipAddress := r.Header.Get("X-Real-Ip")
		userAgent := r.Header.Get("User-Agent")
		password, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		if err != nil {
			return token.TokenDTO{}, err
		}
		err = s.userRepository.Insert(userDTO.UserDTO{
			Name:             dto.Name,
			Password:         string(password),
			Email:            "",
			Role:             role.ROLEUSER,
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

		User, err2 := s.userRepository.GetUserByCredentials(dto.Name, dto.Password)
		if err2 != nil {
			return token.TokenDTO{
				AccessToken:  "",
				RefreshToken: "",
			}, err2
		}
		accessToken, _ := auth.GenerateAccessToken(User.Name)
		refreshToken, err3 := auth.GenerateRefreshToken(User.Name)
		if err3 == nil {
			err4 := s.tokenRepository.CreateToken(refreshToken, User.ID, ipAddress, userAgent)
			if err4 != nil {
				return token.TokenDTO{
					AccessToken:  "",
					RefreshToken: "",
				}, err4
			}
		}
		return token.TokenDTO{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil
	}
	return token.TokenDTO{
		AccessToken:  "",
		RefreshToken: "",
	}, errors.New("not unique username")
}

func (s *AuthService) User(refreshToken string) (request.AuthUserDTO, error) {
	claims, err := auth.ParseToken(refreshToken)
	if err != nil {
		return request.AuthUserDTO{
			Username: "",
			Role:     "",
		}, err
	}
	username := claims["username"].(string)

	User, err2 := s.userRepository.GetUserByUsername(username)
	if err2 != nil {
		return request.AuthUserDTO{
			Username: "",
			Role:     "",
		}, err
	}
	rl := role.RoleDictionary{}.Get(User.Role)

	return request.AuthUserDTO{
		Username: username,
		Role:     rl,
	}, err
}

func (s *AuthService) GetAuthUser(name, password string) (request.AuthUserDTO, error) {
	usr, err := s.userRepository.GetUserByCredentials(name, password)
	if err != nil {
		return request.AuthUserDTO{}, err
	}
	rl := role.RoleDictionary{}.Get(usr.Role)
	var dto = request.AuthUserDTO{
		Username: usr.Name,
		Role:     rl,
	}
	return dto, nil
}
