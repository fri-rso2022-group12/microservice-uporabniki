package main

import (
    docs "microservice-uporabniki/docs"
    "github.com/gin-gonic/gin"
    swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    "microservice-uporabniki/controllers"
    "microservice-uporabniki/initializers"
)

func init(){
    initializers.LoadEnvVariables()
    initializers.ConnectToMysql()
}

func main() {
    r := gin.Default()
    docs.SwaggerInfo.BasePath = ""

    r.POST("/users", controllers.UsersCreate)
    r.GET("/users", controllers.UsersIndex)
    r.GET("/users/:id", controllers.UsersShow)
    r.PUT("/users/:id", controllers.UsersUpdate)
    r.DELETE("/users/:id", controllers.UsersDelete)

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
    r.Run()
}
