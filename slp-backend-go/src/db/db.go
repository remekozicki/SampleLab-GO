package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func InitDB() {
	dsn := "postgres://postgres:postgres@localhost:5432/sample-lab-db"
	var err error
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("❌ Błąd połączenia z bazą: %v", err))
	}
	fmt.Println("✅ Połączono z bazą danych")
}

func SetDB(db *gorm.DB) {
	database = db
}

func GetDB() *gorm.DB {
	return database
}
