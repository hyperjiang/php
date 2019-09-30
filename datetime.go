package php

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Time returns current Unix timestamp
func Time() int64 {
	return time.Now().Unix()
}

// Microtime returns current Unix timestamp with microseconds
func Microtime() float64 {
	return Round(float64(time.Now().UnixNano())/1000000000, 4)
}

// Strtotime parses any English textual datetime description into a Unix timestamp,
// default timezone is UTC if you do not specify one
//
// e.g. You can specify the timezone to UTC+8 by using yyyy-MM-ddThh:mm:ss+08:00
//
// Returns a timestamp on success, -1 on failure.
//
// Note that this function does not fully cover PHP's strtotime function yet
func Strtotime(str string) int64 {
	t, err := DateCreate(str)
	if err != nil {
		return -1
	}
	return t.Unix()
}

// IsLeapYear checks if the given time is in a leap year
func IsLeapYear(t time.Time) bool {
	t2 := time.Date(t.Year(), time.December, 31, 0, 0, 0, 0, time.UTC)
	return t2.YearDay() == 366
}

// LastDateOfMonth gets the last date of the month which the given time is in
func LastDateOfMonth(t time.Time) time.Time {
	t2 := FirstDateOfNextMonth(t)
	return time.Unix(t2.Unix()-86400, 0)
}

// FirstDateOfMonth gets the first date of the month which the given time is in
func FirstDateOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
}

// FirstDateOfNextMonth gets the first date of next month
func FirstDateOfNextMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	if month == time.December {
		year++
		month = time.January
	} else {
		month++
	}
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
}

// FirstDateOfLastMonth gets the first date of last month
func FirstDateOfLastMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	if month == time.January {
		year--
		month = time.December
	} else {
		month--
	}
	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
}

// recognize the character in the php date/time format string
func recognize(c string, t time.Time) string {
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

// parse returns a textual representation of the time value formatted
// according to the given php date/time format string
func parse(format string, t time.Time) string {
	result := ""
	for _, s := range format {
		result += recognize(string(s), t)
	}
	return result
}

// format is a wrapper of parse
func format(f string, t time.Time) string {

	pattern, err := getPattern(f)

	if err != nil {
		return parse(f, t)
	}

	return t.Format(pattern.layout)
}

// Date returns a string formatted according to the given format string using the given integer timestamp
//
// Note that the timezone is set to UTC
func Date(f string, timestamp int64) string {
	return format(f, time.Unix(timestamp, 0).UTC())
}

// Today returns a string formatted according to the given format string using current timestamp
//
// Note that the timezone is using local timezone
func Today(f string) string {
	return format(f, time.Now())
}

// LocalDate returns a string formatted according to the given format string using the given integer timestamp
//
// Note that the timezone is using local timezone
func LocalDate(f string, timestamp int64) string {
	return format(f, time.Unix(timestamp, 0))
}

// Checkdate validates a Gregorian date
func Checkdate(month, day, year int) bool {
	if month < 1 || month > 12 || day < 1 || day > 31 || year <= 0 {
		return false
	}
	_, err := time.Parse("2006-1-2", fmt.Sprintf("%04d-%d-%d", year, month, day))
	return err == nil
}

// DateCreateFromFormat parses a date/time string according to a specified format
func DateCreateFromFormat(f string, t string) (time.Time, error) {
	return time.Parse(convertLayout(f), t)
}

// DateCreate parses a date/time string and return a time.Time.
// Supported Date and Time Formats: https://www.php.net/manual/en/datetime.formats.php
func DateCreate(str string) (time.Time, error) {
	if strings.ToLower(str) == "now" {
		return time.Now(), nil
	}

	reg := regexp.MustCompile("(\\+|\\-)\\s?(\\d)\\s?(day|month|year|week|hour|minute|second)s?")
	matches := reg.FindStringSubmatch(str)
	if matches != nil {
		var diff int64
		num, _ := strconv.ParseInt(matches[2], 10, 64)
		switch matches[3] {
		case "day":
			diff = num * 86400
		case "month":
			diff = num * 86400 * 30
		case "year":
			diff = num * 86400 * 365
		case "week":
			diff = num * 86400 * 7
		case "hour":
			diff = num * 3600
		case "minute":
			diff = num * 60
		case "second":
			diff = num
		}
		if matches[1] == "+" {
			return time.Now().Add(time.Duration(diff) * time.Second), nil
		}
		return time.Now().Add(-time.Duration(diff) * time.Second), nil
	}

	for _, p := range _defaultPatterns {

		reg = regexp.MustCompile(p.regexp)

		if reg.MatchString(str) {
			t, err := time.Parse(p.layout, str)
			if err == nil {
				return t, nil
			}
		}
	}

	return time.Time{}, errors.New("Unsupported date/time string: " + str)
}

// DateDateSet sets the date by year, month and day
func DateDateSet(year, month, day int) (time.Time, error) {
	return time.Parse("2006-1-2", fmt.Sprintf("%04d-%d-%d", year, month, day))
}

// DateDefaultTimezoneGet gets the default timezone
func DateDefaultTimezoneGet() string {
	tz, _ := time.Now().Local().Zone()
	return tz
}

// DateDefaultTimezoneSet sets the default timezone
func DateDefaultTimezoneSet(tz string) error {
	loc, err := time.LoadLocation(tz)
	if err != nil {
		return err
	}
	time.Local = loc
	return nil
}
