package database

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type DB struct {
	dsn string
}

func New() *DB {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatalf("Missing environment variables. Host: %s, Port: %s, User: %s, Password: %s, DB: %s",
			host, port, user, password, dbname)
	}

	cfg := "host=" + host +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbname +
		" port=" + port +
		" sslmode=disable" +
		" connect_timeout=5"

	return &DB{
		dsn: cfg,
	}
}

func (db *DB) Connect() *gorm.DB {
	database, err := gorm.Open(postgres.Open(db.dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to database")
	return database
}
