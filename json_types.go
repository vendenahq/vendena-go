package vendena

import "time"

// String returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

// Time returns a pointer to the time.Time value passed in.
func Time(v time.Time) *time.Time {
	return &v
}

// NullTime returns a pointer to the beginning of time.
func NullTime() *time.Time {
	var t = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)
	return &t
}
