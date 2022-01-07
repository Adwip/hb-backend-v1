package dateTime

import "time"


func DateTimeNow() string{
	currentTime := time.Now()

	result := currentTime.Format("15:04:05 02-01-2006")
	return result
}