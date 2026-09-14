package main

import (
	"lingcard-go/helpers/database"
	"lingcard-go/migrations"
)

func main() {
	db := database.New()
	dbConnect := db.Connect()
	migrator := migrations.NewMigrator(dbConnect)
	migrator.RunMigrations()
}
