package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "../database"
)

func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, "hello world!!!")

  database.ConnectDB()
}
