package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func GetString(key string, defaultValue string) string {
	value := os.Getenv((key))
	if value != "" {
		return value
	}
	return defaultValue
}

func GetInt(key string, defaultValue int) int {
	value := os.Getenv((key))
	if value != "" {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			log.Println(err.Error())
			return defaultValue
		}
		return intValue
	}
	return defaultValue
}
