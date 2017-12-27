package php

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var (
	_now    time.Time
	_zero   time.Time
	_zone   string
	_offset int
)

func init() {
	_now = time.Now()
	_zero = time.Unix(0, 0)
	_zone, _offset = _now.Zone()
}

// Time return current Unix timestamp
func Time() int64 {
	return _now.Unix()
}

// Strtotime parse about any English textual datetime description into a Unix timestamp,
// default timezone is UTC if you did not specify one
//
// e.g. You can specify the timezone to UTC+8 by using yyyy-MM-ddThh:mm:ss+08:00
//
// Returns a timestamp on success, -1 on failure.
//
// Note that this function does not fully cover PHP's strtotime function yet
func Strtotime(str string) int64 {

	if strings.ToLower(str) == "now" {
		return Time()
	}

	for _, p := range _defaultPatterns {

		reg := regexp.MustCompile(p.regexp)

		if reg.MatchString(str) {
			t, err := time.Parse(p.layout, str)
			if err == nil {
				return t.Unix()
			}
		}
	}

	return -1
}

// IsLeapYear check if the given time is in a leap year
func IsLeapYear(t time.Time) bool {
	t2 := time.Date(t.Year(), time.December, 31, 0, 0, 0, 0, time.UTC)
	return t2.YearDay() == 366
}

// LastDateOfMonth get the last date of the month which the given time is in
func LastDateOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	if month == time.December {
		year++
		month = time.January
	} else {
		month++
	}
	t2 := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	return time.Unix(t2.Unix()-86400, 0)
}

// Recognize the character in the format parameter string of Parse
func Recognize(c string, t time.Time) string {
	switch c {
	// Day
	case "d": // Day of the month, 2 digits with leading zeros
		return fmt.Sprintf("%02d", t.Day())
	case "D": // A textual representation of a day, three letters
		return t.Format("Mon")
	case "j": // Day of the month without leading zeros
		return fmt.Sprintf("%d", t.Day())
	case "l": // A full textual representation of the day of the week
		return t.Weekday().String()
	case "w": // Numeric representation of the day of the week
		return fmt.Sprintf("%d", t.Weekday())
	case "z": // The day of the year (starting from 0)
		return fmt.Sprintf("%v", t.YearDay()-1)

	// Week
	case "W": // ISO-8601 week number of year, weeks starting on Monday
		_, w := t.ISOWeek()
		return fmt.Sprintf("%d", w)

	// Month
	case "F": // A full textual representation of a month
		return t.Month().String()
	case "m": // Numeric representation of a month, with leading zeros
		return fmt.Sprintf("%02d", t.Month())
	case "M": // A short textual representation of a month, three letters
		return t.Format("Jan")
	case "n": // Numeric representation of a month, without leading zeros
		return fmt.Sprintf("%d", t.Month())
	case "t": // Number of days in the given month
		return LastDateOfMonth(t).Format("2")

	// Year
	case "L": // Whether it's a leap year
		if IsLeapYear(t) {
			return "1"
		}
		return "0"
	case "o":
		fallthrough
	case "Y": // A full numeric representation of a year, 4 digits
		return fmt.Sprintf("%v", t.Year())
	case "y": // A two digit representation of a year
		return t.Format("06")

	// Time
	case "a": // Lowercase Ante meridiem and Post meridiem
		return t.Format("pm")
	case "A": // Uppercase Ante meridiem and Post meridiem
		return strings.ToUpper(t.Format("pm"))
	case "g": // 12-hour format of an hour without leading zeros
		return t.Format("3")
	case "G": // 24-hour format of an hour without leading zeros
		return fmt.Sprintf("%d", t.Hour())
	case "h": // 12-hour format of an hour with leading zeros
		return t.Format("03")
	case "H": // 24-hour format of an hour with leading zeros
		return fmt.Sprintf("%02d", t.Hour())
	case "i": // Minutes with leading zeros
		return fmt.Sprintf("%02d", t.Minute())
	case "s": // Seconds, with leading zeros
		return fmt.Sprintf("%02d", t.Second())
	case "e": // Timezone identifier
		fallthrough
	case "T": // Timezone abbreviation
		return t.Format("MST")
	case "O": // Difference to Greenwich time (GMT) in hours
		return t.Format("-0700")
	case "P": // Difference to Greenwich time (GMT) with colon between hours and minutes
		return t.Format("-07:00")
	case "U": // Seconds since the Unix Epoch (January 1 1970 00:00:00 GMT)
		return fmt.Sprintf("%v", t.Unix())

	default:
		return c
	}
}

// Parse returns a string formatted according to the given format string using the given time
func Parse(format string, t time.Time) string {
	result := ""
	for _, s := range format {
		result += Recognize(string(s), t)
	}
	return result
}

// Format returns a string formatted according to the given format string using the given time
func Format(format string, t time.Time) string {

	pattern, err := GetPattern(format)

	if err != nil {
		return Parse(format, t)
	}

	return t.Format(pattern.layout)
}

// Date returns a string formatted according to the given format string using the given integer timestamp
//
// Note that the timezone is set to UTC
func Date(format string, timestamp int64) string {
	return Format(format, time.Unix(timestamp, 0).UTC())
}

// Today returns a string formatted according to the given format string using current timestamp
//
// Note that the timezone is using local timezone
func Today(format string) string {
	return Format(format, _now)
}

// LocalDate returns a string formatted according to the given format string using the given integer timestamp
//
// Note that the timezone is using local timezone
func LocalDate(format string, timestamp int64) string {
	return Format(format, time.Unix(timestamp, 0))
}
