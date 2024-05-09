package helpers

import (
	"os"
	"strconv"
)

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetEnvAsInt(key string, defaultValue int) int {
	value := GetEnv(key, "")
	if val, err := strconv.Atoi(value); err != nil {
		return val
	}
	return defaultValue
}
