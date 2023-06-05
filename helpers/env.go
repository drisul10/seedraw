package helpers

import "os"

func GetEnvOrUseDefaultValue(envKey, defaultValue string) string {
	value := os.Getenv(envKey)
	if value == "" {
		value = defaultValue
	}
	return value
}
