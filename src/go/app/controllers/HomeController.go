package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {
	c.JSON(http.StatusOK, "hello home")
}
