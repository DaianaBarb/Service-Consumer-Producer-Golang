package utils

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strconv"
	"time"
)

func PanicRecover() {
	if err := recover(); err != nil {
		log.Fatal(fmt.Sprintf("Panic %s\n%s", err, debug.Stack()), "")
	}
}

func GetTimeNow() *time.Time {
	now := time.Now()
	return &now
}

func GetTimePointer(time time.Time) *time.Time {
	return &time
}

func GetDurationPointer(duration time.Duration) *time.Duration {
	return &duration
}

func GetStringPointer(str string) *string {
	return &str
}

func GetEnvOrFallback(key string, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}

	return value
}

func GetEnvOrFallbackInt32(key string, fallback int32) int32 {
	return int32(GetEnvOrFallbackInt64(key, int64(fallback)))
}

func GetEnvOrFallbackInt64(key string, fallback int64) int64 {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}

	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return fallback
	}

	return i
}

func Contains(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func ContainsValues(valuesSource []string, valuesFind []string) (bool, string) {
	for _, origin := range valuesSource {
		for _, tested := range valuesFind {
			if origin == tested {
				return true, tested
			}
		}
	}
	return false, ""
}

func RemoveIndexInt(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func RemoveIndexString(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func GetTimeFromTimestamp(timestamp string) (*time.Time, error) {
	i, err := strconv.ParseInt(timestamp, 10, 64)

	if err != nil {
		return nil, err
	}

	time := time.Unix(i/1000, 0) // i/1000 convert miliseconds to seconds

	return &time, nil
}

func ParseIntFallbackInt64(value string, fallback int64) int64 {
	if len(value) == 0 {
		return fallback
	}

	i, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return fallback
	}

	return i
}

func ParseIntFallbackInt32(value string, fallback int64) int32 {
	return int32(ParseIntFallbackInt64(value, fallback))
}
