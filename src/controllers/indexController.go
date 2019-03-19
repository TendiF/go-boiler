package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "../database"
  "fmt"
)

func GetIndex(c *gin.Context) {
  fmt.Println("id__")
  database.ConnectDB()
	c.JSON(http.StatusOK, "hello world!!!")
}
