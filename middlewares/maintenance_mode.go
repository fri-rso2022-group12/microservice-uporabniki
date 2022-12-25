package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"microservice-uporabniki/initializers"
	"strings"
)

func MaintenanceMode() gin.HandlerFunc {
	return func(c *gin.Context) {
		pair, _, err := initializers.ConsulKV.Get("MAINTENANCE_MODE", nil)
		if err != nil {
			panic(err)
		}

		if strings.ToLower(string(pair.Value)) == "true" {
			fmt.Println("System in maintenance mode")
			c.AbortWithStatus(503)
		}

	}
}
