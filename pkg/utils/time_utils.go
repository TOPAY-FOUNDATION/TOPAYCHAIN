package utils

import (
	"time"
)

func GetCurrentTimestamp() string {
	return time.Now().Format(time.RFC3339)
}

func ParseTimestamp(timestamp string) (time.Time, error) {
	return time.Parse(time.RFC3339, timestamp)
}

func ElapsedTime(start time.Time) string {
	return time.Since(start).String()
}
