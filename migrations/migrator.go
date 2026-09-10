package migrations

import (
	"fmt"
	"lingcard-go/helpers/database"
	"log"
	"os"
)

func RunMigrations() {
	db := database.New()
	gdb := db.Connect()
	sqlBytes, err := os.ReadFile("migrations/sql/init.sql")
	if err != nil {
		log.Fatalf("read sql/init.sql: %v", err)
	}
	if err := gdb.Exec(string(sqlBytes)).Error; err != nil {
		log.Fatalf("exec init.sql: %v", err)
	}
	fmt.Println("Миграции применены успешно")
}
