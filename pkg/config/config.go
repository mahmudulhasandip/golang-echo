package config

import (
	"echo-auth/pkg/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// db connection
var db *gorm.DB

func Connect() {
	var err error
	dsn := os.Getenv("DB")
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func GetDatabase() *gorm.DB {
	return db
}

func SyncDatabase() {
	if err := db.AutoMigrate(models.User{}); err != nil {
		log.Fatal("Error AutoMigrate: ", err)
		return
	}
}
