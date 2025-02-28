package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIkey string
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
		return nil
	}

	return &Config{
		APIkey: os.Getenv("OWM_API_key"),
	}
}
