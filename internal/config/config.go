package config

import (
	// "log"
	"os"
)

type Config struct {
	OracleHost     string
	OraclePort     string
	OracleUser     string
	OraclePassword string
	OracleService  string
	AppPort        string
}

func LoadConfig() *Config {
	return &Config{
		OracleHost:     getEnv("ORACLE_HOST", "localhost"),
		OraclePort:     getEnv("ORACLE_PORT", "1521"),
		OracleUser:     getEnv("ORACLE_USER", "system"),
		OraclePassword: getEnv("ORACLE_PASSWORD", "admin"),
		OracleService:  getEnv("ORACLE_SERVICE", "XEPDB1"),
		AppPort:        getEnv("APP_PORT", "8080"),
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
