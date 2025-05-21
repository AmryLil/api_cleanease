package utils

import (
	"api_cleanease/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config := config.LoadDBConfig()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		config.DB_HOST, config.DB_USER, config.DB_PASS, config.DB_NAME, config.DB_PORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	} else {
		log.Print("succes connect to db")
	}

	err = migrate(db)
	if err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	return db
}

func migrate(db *gorm.DB) error {

	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	return nil
}
