package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"lingcard-go/handlers/authHandler"
	"lingcard-go/handlers/languageHandler"
	"lingcard-go/handlers/voteHandler"
	"lingcard-go/helpers/database"
	"lingcard-go/repositories/availableLanguageRepository"
	"lingcard-go/repositories/languageRepository"
	"lingcard-go/repositories/tokenRepository"
	"lingcard-go/repositories/userRepository"
	"lingcard-go/repositories/voiceRepository"
	"lingcard-go/repositories/voteOptionRepository"
	"lingcard-go/repositories/voteRepository"
	"lingcard-go/services/authService"
	"lingcard-go/services/languageService"
	"lingcard-go/services/voteService"
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
	voteRepo := voteRepository.New(dbConnect)
	voiceRepo := voiceRepository.New(dbConnect)
	voteOptionRepo := voteOptionRepository.New(dbConnect)

	authSer := authService.New(userRepo, tokenRepo)
	langSer := languageService.New(langRepo, availableLangRepo)
	voteSer := voteService.New(voteRepo, voiceRepo, voteOptionRepo)

	authHand := authHandler.New(authSer)
	langHand := languageHandler.New(langSer)
	voteHand := voteHandler.New(voteSer)

	router.Post("/api/login", authHand.Login)
	router.Post("/api/logout", authHand.Logout)
	router.Post("/api/refresh", authHand.Refresh)
	router.Post("/api/register", authHand.Register)
	router.Post("/api/user", authHand.User)

	router.Get("/api/languages", langHand.GetAllLanguages)
	router.Get("/api/active-languages", langHand.GetAllActiveLanguages)
	router.Get("/api/except-languages/{id}", langHand.GetExceptLanguages)
	router.Get("/api/map-languages", langHand.GetLanguageMap)

	router.Get("/api/votes", voteHand.GetAllVotes)
	router.Get("/api/votes/{id}", voteHand.GetOne)
	router.Post("/api/voices/{voteOptionId}", voteHand.GetAllVotes)
	router.Post("/api/cancel-voices/{voteOptionId}", voteHand.GetAllVotes)

	http.ListenAndServe(":6611", router)
}
