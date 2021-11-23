package config

import (
	"os"
	"strconv"
)

type Configuration struct {
	ServerPort       string
	DatabaseHost     string
	DatabasePort     int
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
}

func SetUpConfiguration() *Configuration {
	config := &Configuration{
		ServerPort:       getEnvStr("VIBEUS_SERVER_LISTEN_ADDR", "8090"),
		DatabaseHost:     getEnvStr("DATABASE_HOST", "database"),
		DatabasePort:     getEnvInt("DATABASE_PORT", 5432),
		DatabaseName:     getEnvStr("APP_DB_NAME", "vibeus"),
		DatabaseUser:     getEnvStr("APP_DB_USER", "vibeus"),
		DatabasePassword: getEnvStr("APP_DB_PASS", "vibeus"),
	}
	return config
}

func getEnvStr(variable string, defaultValue string) string {
	if env, ok := os.LookupEnv(variable); ok {
		return env
	}
	return defaultValue
}

func getEnvInt(variable string, defaultValue int) int {
	valStr, ok := os.LookupEnv(variable)
	if !ok {
		return defaultValue
	}
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}

	return defaultValue
}
