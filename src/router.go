package main

import (
  "./controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", controllers.GetIndex)
	router.GET("/home", controllers.GetHome)
	router.Run(":4000")
}
