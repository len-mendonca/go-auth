package models

import (
	"fmt"

	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func InitDB() {

	connectionString := fmt.Sprintf("user=postgres.ivoljufuybafuicehwln password=%s host=aws-0-ap-south-1.pooler.supabase.com port=5432 dbname=postgres", os.Getenv("DB_PASS"))

	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)

	}

	if err := DB.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	fmt.Println("Postgres connection successful")
}
