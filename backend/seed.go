package main

import (
	"lingcard-go/helpers/database"
	"lingcard-go/migrations"
)

func main() {
	db := database.New()
	dbConnect := db.Connect()
	seed := migrations.NewSeeder(dbConnect)
	seed.RunSeeders()
}
