package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"lingcard-go/helpers/database"
	"net/http"
)

func main() {
	db := database.New()
	db.Connect()
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	http.ListenAndServe(":6611", router)
}
