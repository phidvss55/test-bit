package env

import (
	"fmt"
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
	valueStr, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	valAsInt, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Printf("Error converting %s to int: %v\n", key, err)
		return defaultValue
	}

	return valAsInt
}
