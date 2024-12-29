package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	SECRET	string
	DB_USER	string
	DB_PASS	string
	DB_HOST string
	DB_PORT string
	DB		string
	PORT	string
}


func	NewConfig() *Config{
	err := godotenv.Load(".env")

	var config Config
	if err != nil{
		log.Fatal("Can not Load ENV")
	}

	config.SECRET = os.Getenv("SECRET")
	config.DB_USER = os.Getenv("DB_USER")
	config.DB_PASS = os.Getenv("DB_PASS")
	config.DB_HOST = os.Getenv("DB_HOST")
	config.DB_PORT = os.Getenv("DB_PORT")
	config.DB = os.Getenv("DB")
	config.PORT = os.Getenv("PORT")
	return &config
}
