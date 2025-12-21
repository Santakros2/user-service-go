package config

import (
	// "log"
	"os"
)

type ConfigOracle struct {
	OracleHost     string
	OraclePort     string
	OracleUser     string
	OraclePassword string
	OracleService  string
	AppPort        string
}

type ConfigMySQL struct {
	MySqlUser     string
	MySqlPassword string
	MySqlHost     string
	MySqlPort     string
	MySqlProtocol string
	MySqlDB       string
	AppPort       string
}

func LoadConfigOracle() *ConfigOracle {
	return &ConfigOracle{
		OracleHost:     getEnv("ORACLE_HOST", "localhost"),
		OraclePort:     getEnv("ORACLE_PORT", "1521"),
		OracleUser:     getEnv("ORACLE_USER", "system"),
		OraclePassword: getEnv("ORACLE_PASSWORD", "admin"),
		OracleService:  getEnv("ORACLE_SERVICE", "XEPDB1"),
		AppPort:        getEnv("APP_PORT", "8080"),
	}
}

func LoadConfigMySQL() *ConfigMySQL {
	return &ConfigMySQL{
		MySqlUser:     getEnv("MYSQL_USER", "root"),
		MySqlPassword: getEnv("MYSQL_PASS", "root"),
		MySqlHost:     getEnv("MYSQL_HOST", "localhost"),
		MySqlPort:     getEnv("MYSQL_PORT", "3306"),
		MySqlProtocol: getEnv("MYSQL_PROTOCOL", "tcp"),
		MySqlDB:       getEnv("MYSQL_DB", "userdb"),
		AppPort:       getEnv("APP_PORT", "8080"),
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
