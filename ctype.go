package php

import "regexp"

// CtypeAlnum returns TRUE if every character in text is either a letter or a digit, FALSE otherwise
func CtypeAlnum(text string) bool {
	reg := regexp.MustCompile("^[[:alnum:]]+$") // same pattern ^[a-zA-Z0-9]+$
	return reg.MatchString(text)
}

// CtypeAlpha returns TRUE if every character in text is a letter, FALSE otherwise
func CtypeAlpha(text string) bool {
	reg := regexp.MustCompile("^[[:alpha:]]+$") // same pattern ^[A-Za-z]+$
	return reg.MatchString(text)
}

// CtypeCntrl returns TRUE if every character in text is a control character from the current locale, FALSE otherwise
func CtypeCntrl(text string) bool {
	reg := regexp.MustCompile("^[[:cntrl:]]+$") // same pattern ^[\x00-\x1f\x7f]+$
	return reg.MatchString(text)
}

// CtypeDigit returns TRUE if every character in the string text is a decimal digit, FALSE otherwise
func CtypeDigit(text string) bool {
	reg := regexp.MustCompile("^[[:digit:]]+$") // same pattern ^[0-9]+$
	return reg.MatchString(text)
}

// CtypeGraph returns TRUE if every character in text is printable and actually creates visible output (no white space), FALSE otherwise
func CtypeGraph(text string) bool {
	reg := regexp.MustCompile("^[[:graph:]]+$") // same pattern ^[!-~]+$
	return reg.MatchString(text)
}

// CtypeLower returns TRUE if every character in text is a lowercase letter
func CtypeLower(text string) bool {
	reg := regexp.MustCompile("^[[:lower:]]+$") // same pattern ^[a-z]+$
	return reg.MatchString(text)
}

// CtypePrint returns TRUE if every character in text will actually create output (including blanks),
// returns FALSE if text contains control characters or characters that do not have any output
// or control function at all
func CtypePrint(text string) bool {
	reg := regexp.MustCompile("^[[:print:]]+$") // same pattern ^[ -~]+$
	return reg.MatchString(text)
}

// CtypePunct returns TRUE if every character in text is printable,
// but neither letter, digit or blank, FALSE otherwise
func CtypePunct(text string) bool {
	reg := regexp.MustCompile("^[[:punct:]]+$") // same pattern ^[!-\\/\\:-@\\[-`\\{-~]+$
	return reg.MatchString(text)
}

// CtypeSpace returns TRUE if every character in text creates some sort of white space, FALSE otherwise
//
// Besides the blank character this also includes tab, vertical tab, line feed, carriage return
// and form feed characters
func CtypeSpace(text string) bool {
	reg := regexp.MustCompile("^[[:space:]]+$") // same pattern ^[\\s]+$
	return reg.MatchString(text)
}

// CtypeUpper returns TRUE if every character in text is an uppercase letter
func CtypeUpper(text string) bool {
	reg := regexp.MustCompile("^[[:upper:]]+$") // same pattern ^[A-Z]+$
	return reg.MatchString(text)
}

// CtypeXdigit returns TRUE if every character in text is a hexadecimal 'digit',
// that is a decimal digit or a character from [A-Fa-f] , FALSE otherwise
func CtypeXdigit(text string) bool {
	reg := regexp.MustCompile("^[[:xdigit:]]+$") // same pattern ^[A-Fa-f0-9]+$
	return reg.MatchString(text)
}
