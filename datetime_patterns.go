package php

import (
	"errors"
	"time"
)

// Pattern is a mapping of golang datetime layout and php datetime format
type Pattern struct {
	regexp string // golang regular expression
	layout string // golang datetime layout
	format string // php datetime format
}

// Patterns is an array of Pattern
type Patterns []Pattern

var _defaultPatterns Patterns

func init() {
	_defaultPatterns = Patterns{
		Pattern{
			regexp: "^\\d{4}-\\d{2}-\\d{2}$",
			layout: "2006-01-02",
			format: "Y-m-d",
		},
		Pattern{
			regexp: "^\\d{4}-\\d{1}-\\d{2}$",
			layout: "2006-1-02",
			format: "Y-n-d",
		},
		Pattern{
			regexp: "^\\d{4}-\\d{1}-\\d{1,2}$",
			layout: "2006-1-2",
			format: "Y-n-j",
		},
		Pattern{
			regexp: "^\\d{2}-\\d{2}-\\d{2}$",
			layout: "06-01-02",
			format: "y-m-d",
		},
		Pattern{
			regexp: "^\\d{2}-\\d{1}-\\d{2}$",
			layout: "06-1-02",
			format: "y-n-d",
		},
		Pattern{
			regexp: "^\\d{2}-\\d{1}-\\d{1,2}$",
			layout: "06-1-2",
			format: "y-n-j",
		},
		Pattern{
			regexp: "^\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}$",
			layout: "2006-01-02 15:04:05",
			format: "Y-m-d H:i:s",
		},
		Pattern{
			regexp: "^\\d{8}$",
			layout: "20060102",
			format: "Ymd",
		},
		Pattern{
			regexp: "^\\d{14}$",
			layout: "20060102150405",
			format: "YmdHis",
		},
		Pattern{
			regexp: "^\\d{4}$",
			layout: "2006",
			format: "Y",
		},
		Pattern{
			regexp: "^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}[+-]\\d{2}:\\d{2}$",
			layout: "2006-01-02T15:04:05-07:00",
			format: "c", // ISO 8601 date
		},
		Pattern{
			regexp: "^[[:alpha:]]{3}, \\d{2} [[:alpha:]]{3} \\d{4} \\d{2}:\\d{2}:\\d{2} [+-]\\d{4}$",
			layout: time.RFC1123Z,
			format: "r", // RFC 2822 formatted date
		},
	}

}

// GetPattern get pattern by php style date/time format string
func GetPattern(format string) (Pattern, error) {

	for _, p := range _defaultPatterns {
		if p.format == format {
			return p, nil
		}
	}

	return Pattern{}, errors.New("No pattern found")
}
