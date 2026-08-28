package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"lingcard-go/handlers/authHandler"
	"lingcard-go/handlers/dictionaryHandler"
	"lingcard-go/handlers/languageHandler"
	"lingcard-go/handlers/suggestionHandler"
	"lingcard-go/handlers/voteHandler"
	"lingcard-go/helpers/database"
	"lingcard-go/middlewares/authMiddleware"
	"lingcard-go/repositories/availableLanguageRepository"
	"lingcard-go/repositories/languageRepository"
	"lingcard-go/repositories/suggestionRepository"
	"lingcard-go/repositories/tokenRepository"
	"lingcard-go/repositories/userRepository"
	"lingcard-go/repositories/voiceRepository"
	"lingcard-go/repositories/voteOptionRepository"
	"lingcard-go/repositories/voteRepository"
	"lingcard-go/repositories/wordTranslationRepository"
	"lingcard-go/services/authService"
	"lingcard-go/services/languageService"
	"lingcard-go/services/suggestionService"
	"lingcard-go/services/voteService"
	"lingcard-go/services/wordTranslationService"
	"net/http"
)

func main() {
	db := database.New()
	dbConnect := db.Connect()
	authMid := authMiddleware.New(dbConnect)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	userRepo := userRepository.New(dbConnect)
	tokenRepo := tokenRepository.New(dbConnect)
	langRepo := languageRepository.New(dbConnect)
	availableLangRepo := availableLanguageRepository.New(dbConnect)
	voteRepo := voteRepository.New(dbConnect)
	voiceRepo := voiceRepository.New(dbConnect)
	voteOptionRepo := voteOptionRepository.New(dbConnect)
	wordTranslationRepo := wordTranslationRepository.New(dbConnect)
	suggestionRepo := suggestionRepository.New(dbConnect)

	authSer := authService.New(userRepo, tokenRepo)
	langSer := languageService.New(langRepo, availableLangRepo)
	voteSer := voteService.New(voteRepo, voiceRepo, voteOptionRepo)
	wordTransSer := wordTranslationService.New(wordTranslationRepo)
	suggestionSer := suggestionService.New(suggestionRepo)

	authHand := authHandler.New(authSer)
	langHand := languageHandler.New(langSer)
	voteHand := voteHandler.New(voteSer)
	dictionaryHand := dictionaryHandler.New(wordTransSer)
	suggestionHand := suggestionHandler.New(suggestionSer)

	router.Route("/api", func(r chi.Router) {
		router.Post("/login", authHand.Login)
		router.Post("/logout", authHand.Logout)
		router.Post("/refresh", authHand.Refresh)
		router.Post("/register", authHand.Register)
		router.Post("/user", authHand.User)
	})

	router.Route("/api", func(r chi.Router) {
		r.Use(authMid.Handler)
		r.Get("/languages", langHand.GetAllLanguages)
		r.Get("/active-languages", langHand.GetAllActiveLanguages)
		r.Get("/except-languages/{id}", langHand.GetExceptLanguages)
		r.Get("/map-languages", langHand.GetLanguageMap)
		r.Get("/votes", voteHand.GetAllVotes)
		r.Get("/votes/{id}", voteHand.GetOne)
		r.Post("/voices/{voteOptionId}", voteHand.GetAllVotes)
		r.Post("/cancel-voices/{voteOptionId}", voteHand.GetAllVotes)
		r.Get("/dictionary/{baseLanguageId}/language/{targetLanguageId}", dictionaryHand.Translate)
		r.Post("/suggestions", suggestionHand.Create)
	})
	http.ListenAndServe(":6611", router)
}
