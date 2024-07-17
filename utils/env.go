package utils

import (
	"os"
	"strconv"
)

func GetEnvDefault[T any](key string, defaultValue T) T {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	var result T

	switch any(result).(type) {
	case int:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue
		}
		return any(intValue).(T)
	case float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return defaultValue
		}
		return any(floatValue).(T)
	case bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return defaultValue
		}
		return any(boolValue).(T)
	case string:
		return any(value).(T)
	default:
		return defaultValue
	}
}
