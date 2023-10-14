package exercism

import (
	"fmt"
	"time"
)

// parsingTime parses a string date using the given layout.
func parsingTime(layout, date string) time.Time {
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	return t
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	timeParsed := parsingTime(layout, date)
	return timeParsed
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	t := parsingTime(layout, date)
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t := parsingTime(layout, date)
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t := Schedule(date)
	formatedTime := t.Format("Monday, January 2, 2006, at 15:04")
	return fmt.Sprintf("You have an appointment on %s.", formatedTime)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	layout := "2006-01-02 15:04:05"
	anniversary := parsingTime(layout, "2020-09-15 00:00:00")
	YearsBetween := time.Now().Year() - anniversary.Year()

	anniversaryThisYear := anniversary.AddDate(YearsBetween, 0, 0)
	return anniversaryThisYear
}
