package config

import "os"

type Config struct {
	ServerPort string
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

var BackendConfig Config

func LoadConfig() {
	BackendConfig = Config{
		ServerPort: getEnv("SERVER_PORT", "5000"),
		DBHost: getEnv("DB_HOST", "postgres-db"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBUser: getEnv("DB_USER", "user"),
		DBPass: getEnv("DB_PASS", "password"),
		DBName: getEnv("DB_NAME", "sensor_db"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}