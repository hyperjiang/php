package php

import (
	"errors"
	"time"
)

// pattern stores the mapping of golang datetime layout and php datetime format
type pattern struct {
	regexp string // golang regular expression
	layout string // golang datetime layout
	format string // php datetime format
}

// patterns is an array of pattern
type patterns []pattern

var _defaultPatterns patterns

// formatMap is the mapping of php date format to golang layout
var formatMap map[string]string

func init() {
	_defaultPatterns = patterns{
		pattern{
			regexp: "^\\d{4}-\\d{2}-\\d{2}$",
			layout: "2006-01-02",
			format: "Y-m-d",
		},
		pattern{
			regexp: "^\\d{4}-\\d{2}-\\d{1}$",
			layout: "2006-01-2",
			format: "Y-m-j",
		},
		pattern{
			regexp: "^\\d{4}-\\d{1}-\\d{2}$",
			layout: "2006-1-02",
			format: "Y-n-d",
		},
		pattern{
			regexp: "^\\d{4}-\\d{1}-\\d{1,2}$",
			layout: "2006-1-2",
			format: "Y-n-j",
		},
		pattern{
			regexp: "^\\d{2}-\\d{2}-\\d{2}$",
			layout: "06-01-02",
			format: "y-m-d",
		},
		pattern{
			regexp: "^\\d{2}-\\d{2}-\\d{1}$",
			layout: "06-01-2",
			format: "y-m-j",
		},
		pattern{
			regexp: "^\\d{2}-\\d{1}-\\d{2}$",
			layout: "06-1-02",
			format: "y-n-d",
		},
		pattern{
			regexp: "^\\d{2}-\\d{1}-\\d{1,2}$",
			layout: "06-1-2",
			format: "y-n-j",
		},
		pattern{
			regexp: "^\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}$",
			layout: "2006-01-02 15:04:05",
			format: "Y-m-d H:i:s",
		},
		pattern{
			regexp: "^\\d{4}-\\d{2}-\\d{1} \\d{2}:\\d{2}:\\d{2}$",
			layout: "2006-01-2 15:04:05",
			format: "Y-m-j H:i:s",
		},
		pattern{
			regexp: "^\\d{4}-\\d{1}-\\d{2} \\d{2}:\\d{2}:\\d{2}$",
			layout: "2006-1-02 15:04:05",
			format: "Y-n-d H:i:s",
		},
		pattern{
			regexp: "^\\d{4}-\\d{1}-\\d{1} \\d{2}:\\d{2}:\\d{2}$",
			layout: "2006-1-2 15:04:05",
			format: "Y-n-j H:i:s",
		},
		pattern{
			regexp: "^\\d{2}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}$",
			layout: "06-01-02 15:04:05",
			format: "y-m-d H:i:s",
		},
		pattern{
			regexp: "^\\d{2}-\\d{2}-\\d{1} \\d{2}:\\d{2}:\\d{2}$",
			layout: "06-01-2 15:04:05",
			format: "y-m-j H:i:s",
		},
		pattern{
			regexp: "^\\d{2}-\\d{1}-\\d{2} \\d{2}:\\d{2}:\\d{2}$",
			layout: "06-1-02 15:04:05",
			format: "y-n-d H:i:s",
		},
		pattern{
			regexp: "^\\d{2}-\\d{1}-\\d{1} \\d{2}:\\d{2}:\\d{2}$",
			layout: "06-1-2 15:04:05",
			format: "y-n-j H:i:s",
		},
		pattern{
			regexp: "^\\d{8}$",
			layout: "20060102",
			format: "Ymd",
		},
		pattern{
			regexp: "^\\d{14}$",
			layout: "20060102150405",
			format: "YmdHis",
		},
		pattern{
			regexp: "^\\d{4}$",
			layout: "2006",
			format: "Y",
		},
		pattern{
			regexp: "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}[+-]\\d{2}:\\d{2}$",
			layout: "2006-01-02T15:04:05-07:00",
			format: "c", // ISO 8601 date
		},
		pattern{
			regexp: "^[[:alpha:]]{3}, \\d{2} [[:alpha:]]{3} \\d{4} \\d{2}:\\d{2}:\\d{2} [+-]\\d{4}$",
			layout: time.RFC1123Z,
			format: "r", // RFC 2822 formatted date
		},
		pattern{
			regexp: "^\\d{2} [[:alpha:]]{3} \\d{4}$",
			layout: "02 Jan 2006",
			format: "d M Y",
		},
		pattern{
			regexp: "^\\d{2} [[:alpha:]]{3} \\d{2}$",
			layout: "02 Jan 06",
			format: "d M y",
		},
		pattern{
			regexp: "^\\d{1} [[:alpha:]]{3} \\d{4}$",
			layout: "2 Jan 2006",
			format: "j M Y",
		},
		pattern{
			regexp: "^\\d{1} [[:alpha:]]{3} \\d{2}$",
			layout: "2 Jan 06",
			format: "j M y",
		},
		pattern{
			regexp: "^[[:alpha:]]{3} \\d{2} \\d{4}$",
			layout: "Jan 02 2006",
			format: "M d Y",
		},
		pattern{
			regexp: "^[[:alpha:]]{3} \\d{2} \\d{2}$",
			layout: "Jan 02 06",
			format: "M d y",
		},
		pattern{
			regexp: "^[[:alpha:]]{3} \\d{1} \\d{4}$",
			layout: "Jan 2 2006",
			format: "M j Y",
		},
		pattern{
			regexp: "^[[:alpha:]]{3} \\d{1} \\d{2}$",
			layout: "Jan 2 06",
			format: "M j y",
		},
	}

	formatMap = map[string]string{
		"d": "02",
		"j": "2",
		"m": "01",
		"n": "1",
		"Y": "2006",
		"y": "06",
		"H": "15",
		"h": "03",
		"g": "3",
		"i": "04",
		"s": "05",
		"M": "Jan",
		"F": "January",
		"D": "Mon",
		"l": "Monday",
		"e": "MST",
		"T": "MST",
		"O": "-0700",
		"P": "-07:00",
	}
}

// getPattern gets the matched pattern by the given php style date/time format string
func getPattern(format string) (pattern, error) {
	for _, p := range _defaultPatterns {
		if p.format == format {
			return p, nil
		}
	}

	return pattern{}, errors.New("No pattern found")
}

// convertLayout converts php date format string to golang date layout
func convertLayout(format string) string {
	var layout string
	for _, s := range format {
		if v, ok := formatMap[string(s)]; ok {
			layout += v
		} else {
			layout += string(s)
		}
	}
	return layout
}
