package authHandler

import (
	"encoding/json"
	"io"
	"lingcard-go/dto/request"
	"lingcard-go/services/authService"
	"net/http"
	"os"
	"strconv"
	"time"
)

type AuthHandler struct {
	authService *authService.AuthService
}

func New(authService *authService.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var accessTime, _ = strconv.Atoi(os.Getenv("ACCESS_TOKEN_TIME_EXPIRE"))
	var refreshTime, _ = strconv.Atoi(os.Getenv("REFRESH_TOKEN_TIME_EXPIRE"))
	var credentials request.LoginDTO

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokens, err := h.authService.Login(credentials.Name, credentials.Password, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	accessCookie := &http.Cookie{
		Name:     "access_token",
		Value:    tokens.AccessToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(accessTime) * time.Minute),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	refreshCookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(refreshTime) * time.Minute),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, accessCookie)
	http.SetCookie(w, refreshCookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	refreshCookie, err := r.Cookie("refresh_token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	_ = h.authService.Logout(refreshCookie.Value)

	accessCookie := &http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(0) * time.Minute),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	refreshCookie = &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(0) * time.Minute),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, accessCookie)
	http.SetCookie(w, refreshCookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var accessTime, _ = strconv.Atoi(os.Getenv("ACCESS_TOKEN_TIME_EXPIRE"))
	var refreshTime, _ = strconv.Atoi(os.Getenv("REFRESH_TOKEN_TIME_EXPIRE"))
	var credentials request.RegisterDTO

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokens, err := h.authService.Register(credentials, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	accessCookie := &http.Cookie{
		Name:     "access_token",
		Value:    tokens.AccessToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(accessTime) * time.Minute),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	refreshCookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(refreshTime) * time.Minute),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, accessCookie)
	http.SetCookie(w, refreshCookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var accessTime, _ = strconv.Atoi(os.Getenv("ACCESS_TOKEN_TIME_EXPIRE"))
	var refreshTime, _ = strconv.Atoi(os.Getenv("REFRESH_TOKEN_TIME_EXPIRE"))
	refreshCookie, err := r.Cookie("refresh_token")
	if err != nil || refreshCookie.Value == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokens, err := h.authService.Refresh(refreshCookie.Value, r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	accessCookie := &http.Cookie{
		Name:     "access_token",
		Value:    tokens.AccessToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(accessTime) * time.Minute),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	refreshCookie = &http.Cookie{
		Name:     "refresh_token",
		Value:    tokens.RefreshToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Duration(refreshTime) * time.Minute),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, accessCookie)
	http.SetCookie(w, refreshCookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) User(w http.ResponseWriter, r *http.Request) {
	accessCookie, err := r.Cookie("access_token")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	accessToken := accessCookie.Value
	if accessToken == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	authUserDTO, err := h.authService.User(accessToken)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(authUserDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
