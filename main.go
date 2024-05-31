package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/len-mendonca/go-auth/models"
	"github.com/len-mendonca/go-auth/router"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	models.InitDB()
}

func main() {

	r := gin.Default()

	router.AuthRoutes(r)

	r.Run(":8080")
}
