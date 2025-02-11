package utils

import (
	"api_cleanease/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	config := config.LoadDBConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USER, config.DB_PASS, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Printf("Connecting to DB at %s:%s...\n", config.DB_HOST, config.DB_PORT)

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
