package db

import (
	"fmt"
	"gorm.io/gorm/schema"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var testDB *gorm.DB

func InitDB() {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5433"
	}
	user := os.Getenv("DB_USERNAME")
	if user == "" {
		user = "postgres"
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "postgres"
	}
	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "sample-lab-db"
	}

	fmt.Printf("Connecting to DB host=%s port=%s\n", host, port)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	var err error
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // <- KLUCZOWA LINIA
		},
	})
	if err != nil {
		panic(fmt.Sprintf("Błąd połączenia z bazą: %v", err))
	}
	fmt.Println("Połączono z bazą danych")
}

func SetDB(db *gorm.DB) {
	database = db
}

func GetDB() *gorm.DB {
	if testDB != nil {
		return testDB
	}

	return database
}

func OverrideDB(mock *gorm.DB) {
	testDB = mock
}
