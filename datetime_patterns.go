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

func init() {
	_defaultPatterns = patterns{
		pattern{
			regexp: "^\\d{4}-\\d{2}-\\d{2}$",
			layout: "2006-01-02",
			format: "Y-m-d",
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
