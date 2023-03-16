package utils

import "time"

func GetCurrentDateTimeUTC() time.Time {
	return time.Now().UTC()
}
