package main

import (
  "./controllers"
  "./controllers/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", controllers.GetIndex)
	router.GET("/home", controllers.GetHome)
	router.POST("/register", authController.RegisterUser)
	router.Run(":3000")
}
