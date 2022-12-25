package controllers

import (
	"github.com/gin-gonic/gin"
	"microservice-uporabniki/initializers"
	"net/http"
)

func Health(c *gin.Context) {
	err := initializers.GetDb().Ping()
	if err != nil {
		c.AbortWithStatus(500)
	}
	c.Status(http.StatusOK)
}
