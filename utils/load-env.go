package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv(env string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Load Env Fail")
	}

	return os.Getenv(env)
}
