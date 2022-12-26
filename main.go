package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/heptiolabs/healthcheck"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"microservice-uporabniki/controllers"
	docs "microservice-uporabniki/docs"
	"microservice-uporabniki/initializers"
	"microservice-uporabniki/middlewares"
	"time"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToMysql()
	initializers.InitializeConsul()
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(middlewares.MaintenanceMode())
	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	docs.SwaggerInfo.BasePath = ""

	health := healthcheck.NewHandler()
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))
	health.AddReadinessCheck(
		"database_check",
		healthcheck.DatabasePingCheck(initializers.GetDb(), 100*time.Millisecond))

	r.POST("/users", controllers.UsersCreate)
	r.GET("/users", controllers.UsersIndex)
	r.GET("/users/:id", controllers.UsersShow)
	r.PUT("/users/:id", controllers.UsersUpdate)
	r.DELETE("/users/:id", controllers.UsersDelete)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/healthz", controllers.Health)
	r.Run()
}
