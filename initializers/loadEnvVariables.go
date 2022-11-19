package initializers

import (
    "github.com/joho/godotenv"
    "log"
)

func LoadEnvVariables(){
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file (" + err.Error() + ") - fallbacking to environmental variables")
    }
}
