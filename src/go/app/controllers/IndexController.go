package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	c.JSON(http.StatusOK, "hello world!")
}
