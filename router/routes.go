package router

import "github.com/gin-gonic/gin"

func AuthRoutes(r *gin.Engine) {
	r.POST("/signup")
	r.POST("/login")
	r.GET("/home")
	r.GET("/logout")
	r.GET("/admin")

}
