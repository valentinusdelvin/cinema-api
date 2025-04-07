package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnvVariables() {
	fmt.Println("Running LoadEnvVariables, STAGE =", os.Getenv("STAGE"))
	if os.Getenv("STAGE") == "PROD" {
		return
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
