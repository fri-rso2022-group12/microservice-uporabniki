package main

import (
    "microservice-uporabniki/initializers"
    "microservice-uporabniki/models"
)

func init(){
    initializers.LoadEnvVariables()
    initializers.ConnectToMysql()
}

func main(){
    initializers.DB.AutoMigrate(&models.User{})
}
