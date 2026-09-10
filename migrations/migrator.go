package migrations

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"lingcard-go/helpers/database"
	"log"
	"os"
)

type Migrator struct {
	db *gorm.DB
}

func NewMigrator(db *gorm.DB) *Migrator {
	return &Migrator{
		db: db,
	}
}

func (m *Migrator) RunMigrations() {
	db := database.New()
	gdb := db.Connect()
	sqlBytes, err := os.ReadFile("migrations/sql/init.sql")
	if err != nil {
		log.Fatalf("read sql/init.sql: %v", err)
	}
	if err := gdb.Exec(string(sqlBytes)).Error; err != nil {
		log.Fatalf("exec init.sql: %v", err)
	}
	dbs, _ := gdb.DB()
	defer func(dbs *sql.DB) {
		err := dbs.Close()
		if err != nil {
			log.Fatalf("close database: %v", err)
		}
	}(dbs)
	fmt.Println("Миграции применены успешно")
}
