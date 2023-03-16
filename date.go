package meUtils

import "time"

func GetCurrentDateTimeUTC() time.Time {
	return time.Now().UTC()
}
