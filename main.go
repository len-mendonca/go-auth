package main

import (
	"github.com/gin-gonic/gin"
	"github.com/len-mendonca/go-auth/models"
	"github.com/len-mendonca/go-auth/router"
)

func init() {
	models.InitDB()
}

func main() {
	r := gin.Default()

	router.AuthRoutes(r)

	r.Run(":8080")
}
