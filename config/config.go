package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func GetConfig() Config {
	return Config{
		DBHost: getEnvOrDefault("DB_HOST", ""),
		DBPort: getEnvOrDefault("DB_PORT", ""),
		DBUser: getEnvOrDefault("DB_USER", ""),
		DBPass: getEnvOrDefault("DB_PASSWORD", ""),
		DBName: getEnvOrDefault("DB_NAME", ""),
	}
}

func GetConectionString() string {
	config := GetConfig()

	// Verificar que la contraseña esté configurada
	if config.DBPass == "" {
		panic("DB_PASSWORD no está configurada en las variables de entorno")
	}

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		config.DBHost,
		config.DBUser,
		config.DBPass,
		config.DBName,
		config.DBPort,
	)
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
