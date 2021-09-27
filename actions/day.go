package actions

import "time"

// IsWeekDay checks if the current day of the week is a weekday
func IsWeekDay() bool {
	dayOfWeek := time.Now().Weekday()

	switch dayOfWeek {
	case time.Saturday:
		return false
	case time.Sunday:
		return false
	default:
		return true
	}
}
