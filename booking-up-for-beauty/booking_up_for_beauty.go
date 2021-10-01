package booking

import (
	"fmt"
	"log"
	"time"
)

var forms = []string{"1/2/2006 15:04:05", "Monday 2, 2006 15:04:05 MST", "January 2, 2006 15:04:05", "Monday, January 2, 2006 15:04:05", "Monday, January 2, 2006", "2006-01-2"}

var doParse = func(date string) time.Time {
	var loc, _ = time.LoadLocation("UTC")
	var t time.Time
	for _, v := range forms {
		log.Println("Time-to-Parse: ", date, " using ", v)
		t, err := time.ParseInLocation(v, date, loc)
		if err == nil {
			log.Println("Parsed time: ", t)
			return t
		}
	}
	log.Println("Parsed time: ", t)
	return t
}

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	t := doParse(date)
	return t
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	var t time.Time = Schedule(date)
	var n = time.Now()
	return n.After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	var t time.Time = Schedule(date)
	log.Println("Hours: ", t.Hour(), "Minute: ", t.Minute())
	if t.Hour() >= 12 && t.Hour() < 18 {
		return true
	}

	return false
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	var t time.Time = Schedule(date)

	return fmt.Sprintf("You have an appointment on %s, at %2d:%2d.", t.Format(forms[4]), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	return doParse("2021-09-15")
}
