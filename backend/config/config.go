package config

import (
	"os"
	"strconv"
)

type Config struct {
	MaxAttempts int

	DBEndpoint  string
	DBUser      string
	DBPassword  string
	DBNamespace string
	DBDatabase  string

	MinIOAccessKeyID string
	MinIOSecretKey   string
	MinIOUseSSL      bool
	MinIOBucketName  string
	MinIOEndpoint    string

	NEXT_PUBLIC_SERVER_IP string
}

func LoadConfig() *Config {
	return &Config{
		MaxAttempts: 5,

		DBEndpoint:  "ws://localhost:8000/rpc",
		DBUser:      "root",
		DBPassword:  "root",
		DBNamespace: "vending",
		DBDatabase:  "vending",

		MinIOAccessKeyID: "minioadmin",
		MinIOSecretKey:   "minioadmin",
		MinIOUseSSL:      false, // Change to "true" if you are using https
		MinIOBucketName:  "vending",
		MinIOEndpoint:    "localhost:9009",

		NEXT_PUBLIC_SERVER_IP: "http://localhost",

		// -------------------------------------------------------

		// MaxAttempts: getIntEnv("MAX_ATTEMPTS", 5),

		// DBEndpoint: os.Getenv("DB_ENDPOINT"),
		// DBUser:     os.Getenv("DB_USER"),
		// DBPassword: os.Getenv("DB_PASSWORD"),

		// NEXT_PUBLIC_SERVER_IP: os.Getenv("NEXT_PUBLIC_SERVER_IP"),
	}
}

func getIntEnv(key string, defaultValue int) int {
	if val, ok := os.LookupEnv(key); ok {
		if intVal, err := strconv.Atoi(val); err == nil {
			return intVal
		}
	}
	return defaultValue
}
