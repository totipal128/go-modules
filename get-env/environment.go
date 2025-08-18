package get_env

import (
	"os"
	"strconv"
	"time"
)

func String(key, defaultVal string) string {
	value, exist := os.LookupEnv(key)
	if !exist {
		return defaultVal
	}
	return value
}

// GetBool for get get-env type bool
func GetBool(key string, defaultVal bool) bool {
	value, exist := os.LookupEnv(key)
	if !exist || value == "" {
		return defaultVal
	}

	b, err := strconv.ParseBool(value)
	if err != nil {
		return defaultVal
	}
	return b
}

// GetTimeDuration for get get-env variables type time duration in golang
func TimeDuration(key string, defaultVal time.Duration) time.Duration {
	if os.Getenv(key) == "" {
		return defaultVal
	}

	i, _ := strconv.ParseInt(os.Getenv(key), 10, 64)
	return time.Duration(i)
}

// GetInt64 for get get-env variables type int64 in golang
func Int64(key string, defaultVal int64) int64 {
	if os.Getenv(key) == "" {
		return defaultVal
	}

	i, _ := strconv.ParseInt(os.Getenv(key), 10, 64)
	return i
}

// GetFloat64 for get get-env variables type int64 in golang
func Float64(key string, defaultVal float64) float64 {
	if os.Getenv(key) == "" {
		return defaultVal
	}

	i, _ := strconv.ParseFloat(os.Getenv(key), 64)
	return i
}
