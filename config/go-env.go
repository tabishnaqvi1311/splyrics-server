package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


func getMongoUri() string{
	err := godotenv.Load()

	if err != nil{
		log.Fatal("Env Err")
	}

	return os.Getenv("MONGO_URI")
}