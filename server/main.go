package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"lingcard-go/handlers/authHandler"
	"lingcard-go/handlers/dictionaryHandler"
	"lingcard-go/handlers/languageHandler"
	"lingcard-go/handlers/postHandler"
	"lingcard-go/handlers/profileHandler"
	"lingcard-go/handlers/progressHandler"
	"lingcard-go/handlers/reactionHandler"
	"lingcard-go/handlers/suggestionHandler"
	"lingcard-go/handlers/voteHandler"
	"lingcard-go/helpers/database"
	"lingcard-go/middlewares/authMiddleware"
	"lingcard-go/repositories/availableLanguageRepository"
	"lingcard-go/repositories/commentRepository"
	"lingcard-go/repositories/courseRepository"
	"lingcard-go/repositories/languageRepository"
	"lingcard-go/repositories/postRepository"
	"lingcard-go/repositories/reactionRepository"
	"lingcard-go/repositories/suggestionRepository"
	"lingcard-go/repositories/tokenRepository"
	"lingcard-go/repositories/userRepository"
	"lingcard-go/repositories/voiceRepository"
	"lingcard-go/repositories/voteOptionRepository"
	"lingcard-go/repositories/voteRepository"
	"lingcard-go/repositories/wordTranslationRepository"
	"lingcard-go/services/authService"
	"lingcard-go/services/courseService"
	"lingcard-go/services/languageService"
	"lingcard-go/services/postService"
	"lingcard-go/services/reactionService"
	"lingcard-go/services/suggestionService"
	"lingcard-go/services/userService"
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
	postRepo := postRepository.New(dbConnect)
	reactionRepo := reactionRepository.New(dbConnect)
	commentRepo := commentRepository.New(dbConnect)
	courseRepo := courseRepository.New(dbConnect)

	authSer := authService.New(userRepo, tokenRepo)
	langSer := languageService.New(langRepo, availableLangRepo)
	voteSer := voteService.New(voteRepo, voiceRepo, voteOptionRepo)
	wordTransSer := wordTranslationService.New(wordTranslationRepo)
	suggestionSer := suggestionService.New(suggestionRepo)
	postSer := postService.New(postRepo, langRepo, reactionRepo, commentRepo, userRepo)
	reactionSer := reactionService.New(reactionRepo, postRepo)
	userSer := userService.New(userRepo, courseRepo, wordTranslationRepo)
	courseSer := courseService.New(courseRepo, wordTranslationRepo)

	authHand := authHandler.New(authSer)
	langHand := languageHandler.New(langSer)
	voteHand := voteHandler.New(voteSer)
	dictionaryHand := dictionaryHandler.New(wordTransSer)
	suggestionHand := suggestionHandler.New(suggestionSer)
	postHand := postHandler.New(postSer)
	reactionHand := reactionHandler.New(reactionSer)
	profileHand := profileHandler.New(userSer, courseSer)
	progressHand := progressHandler.New(courseSer)

	router.Route("/api", func(r chi.Router) {
		r.Post("/login", authHand.Login)
		r.Post("/logout", authHand.Logout)
		r.Post("/refresh", authHand.Refresh)
		r.Post("/register", authHand.Register)
		r.Post("/user", authHand.User)

		r.Get("/posts/{code}", postHand.GetPostsByCode)
		r.Get("/posts", postHand.GetPostsByCode)

		r.Group(func(r chi.Router) {
			r.Use(authMid.Handler)
			r.Get("/articles/{id}", postHand.GetOne)
			r.Post("/posts/{id}/comments", postHand.CreateComment)
			r.Post("/posts", postHand.Create)
			r.Delete("/comments/{id}", postHand.DeleteComment)

			r.Post("/likes/{postId}", reactionHand.Like)
			r.Post("/dislikes/{postId}", reactionHand.Dislike)
			r.Post("/unset-reactions/{postId}", reactionHand.Unset)

			r.Post("/logout", authHand.Logout)
			r.Post("/user", authHand.User)
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

			r.Get("/progress/{status}", progressHand.Progress)
			r.Delete("/progress", progressHand.ClearProgress)
			r.Delete("/words/{id}/progress", progressHand.ClearWordProgress)

			r.Get("/profile", profileHand.Profile)
			r.Patch("/profile", profileHand.Update)
		})

	})
	http.ListenAndServe(":6611", router)
}
