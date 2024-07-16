package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	AUTH_PORT          string
	FORUM_SERVICE_PORT string

	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	MONGO_URI             string
	MONGO_DB_NAME         string
	MONGO_COLLECTION_NAME string

	LOG_PATH string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.AUTH_PORT = cast.ToString(coalesce("AUTH_PORT", ":8088"))
	config.FORUM_SERVICE_PORT = cast.ToString(coalesce("FORUM_SERVICE_PORT", ":50051"))

	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "n10"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "12345"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "n10"))

	config.MONGO_URI = cast.ToString(coalesce("MONGO_URI", "mongodb://localhost:27017"))
	config.MONGO_DB_NAME = cast.ToString(coalesce("MONGO_DB_NAME", "lingua_learning"))
	config.MONGO_COLLECTION_NAME = cast.ToString(coalesce("MONGO_COLLECTION_NAME", "exercises"))

	config.LOG_PATH = cast.ToString(coalesce("LOG_PATH", "logs/info.log"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}