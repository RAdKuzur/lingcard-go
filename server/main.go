package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"lingcard-go/handlers/authHandler"
	"lingcard-go/helpers/database"
	"lingcard-go/repositories/tokenRepository"
	"lingcard-go/repositories/userRepository"
	"lingcard-go/services/authService"
	"net/http"
)

func main() {
	db := database.New()
	dbConnect := db.Connect()
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	userRepo := userRepository.New(dbConnect)
	tokenRepo := tokenRepository.New(dbConnect)

	authSer := authService.New(userRepo, tokenRepo)

	authHand := authHandler.New(authSer)

	router.Post("/api/login", authHand.Login)
	router.Post("/api/logout", authHand.Logout)
	router.Post("/api/refresh", authHand.Refresh)
	router.Post("/api/register", authHand.Register)
	router.Post("/api/user", authHand.User)
	http.ListenAndServe(":6611", router)
}
