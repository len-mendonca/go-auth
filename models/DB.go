package models

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := fmt.Sprintf("user=postgres.ivoljufuybafuicehwln password=%s host=aws-0-ap-south-1.pooler.supabase.com port=5432 dbname=postgres", os.Getenv("DB_PASS"))

	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Postgres connection successful")

}