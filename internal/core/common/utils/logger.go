package utils

import (
	"time"
)

func GetCurrentTime() string {
	currentTime := time.Now()

	timeString := currentTime.Format("2006-01-02 15:04:05")

	return timeString
}

