package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"lingcard-go/handlers/authHandler"
	"lingcard-go/handlers/languageHandler"
	"lingcard-go/helpers/database"
	"lingcard-go/repositories/availableLanguageRepository"
	"lingcard-go/repositories/languageRepository"
	"lingcard-go/repositories/tokenRepository"
	"lingcard-go/repositories/userRepository"
	"lingcard-go/services/authService"
	"lingcard-go/services/languageService"
	"net/http"
)

func main() {
	db := database.New()
	dbConnect := db.Connect()
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	userRepo := userRepository.New(dbConnect)
	tokenRepo := tokenRepository.New(dbConnect)
	langRepo := languageRepository.New(dbConnect)
	availableLangRepo := availableLanguageRepository.New(dbConnect)

	authSer := authService.New(userRepo, tokenRepo)
	langSer := languageService.New(langRepo, availableLangRepo)

	authHand := authHandler.New(authSer)
	langHand := languageHandler.New(langSer)

	router.Post("/api/login", authHand.Login)
	router.Post("/api/logout", authHand.Logout)
	router.Post("/api/refresh", authHand.Refresh)
	router.Post("/api/register", authHand.Register)
	router.Post("/api/user", authHand.User)

	router.Get("/api/languages", langHand.GetAllLanguages)
	router.Get("/api/active-languages", langHand.GetAllActiveLanguages)
	router.Get("/api/except-languages/{id}", langHand.GetExceptLanguages)
	router.Get("/api/map-languages", langHand.GetLanguageMap)

	http.ListenAndServe(":6611", router)
}
