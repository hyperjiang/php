package php

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"hash/crc32"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Substr returns the portion of string specified by the start and length parameters.
//
// The behaviour of this function is mostly the same as the PHP mb_substr function,
//
// see http://php.net/manual/en/function.mb-substr.php
//
// except that:
//
// 1) If start or length is invalid, empty string will be return;
//
// 2) If length is 0, the substring starting from start until the end of the string will be returned.
func Substr(str string, start, length int) string {

	rs := []rune(str)
	rl := len(rs)

	if rl == 0 {
		return ""
	}

	if start < 0 {
		start = rl + start
	}
	if start < 0 {
		start = 0
	}
	if start > rl-1 {
		return ""
	}

	end := rl

	if length < 0 {
		end = rl + length
	} else if length > 0 {
		end = start + length
	}

	if end < 0 || start >= end {
		return ""
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

// Strlen get string length
//
// A multi-byte character is counted as 1
func Strlen(str string) int {
	return len([]rune(str))
}

// Strpos find position of first occurrence of string in a string
//
// It's multi-byte safe. return -1 if can not find the substring
func Strpos(haystack, needle string) int {

	pos := strings.Index(haystack, needle)

	if pos < 0 {
		return pos
	}

	rs := []rune(haystack[0:pos])

	return len(rs)
}

// Strrpos find the position of the last occurrence of a substring in a string
func Strrpos(haystack, needle string) int {

	pos := strings.LastIndex(haystack, needle)

	if pos < 0 {
		return pos
	}

	rs := []rune(haystack[0:pos])

	return len(rs)
}

// Stripos find position of the first occurrence of a case-insensitive substring in a string
func Stripos(haystack, needle string) int {
	return Strpos(strings.ToLower(haystack), strings.ToLower(needle))
}

// Strripos find the position of the last occurrence of a case-insensitive substring in a string
func Strripos(haystack, needle string) int {
	return Strrpos(strings.ToLower(haystack), strings.ToLower(needle))
}

// buildReplaceSlice is a helper function for Replace and Ireplace
func buildReplaceSlice(search, replace interface{}) ([]string, []string, error) {

	var aSearch, aReplace []string

	switch s := search.(type) {
	case int:
		aSearch = append(aSearch, strconv.Itoa(s))
	case rune:
		aSearch = append(aSearch, string(s))
	case string:
		aSearch = append(aSearch, s)
	case []string:
		aSearch = s
	default:
		return aSearch, aReplace, errors.New("unsupported type of search")
	}

	switch r := replace.(type) {
	case int:
		aReplace = append(aReplace, strconv.Itoa(r))
	case rune:
		aReplace = append(aReplace, string(r))
	case string:
		aReplace = append(aReplace, r)
	case []string:
		aReplace = r
	default:
		return aSearch, aReplace, errors.New("unsupported type of replace")
	}

	return aSearch, aReplace, nil
}

// Replace all occurrences of the search string with the replacement string
//
// This function is an implement of PHP's str_replace
//
// see http://php.net/manual/en/function.str-replace.php
func Replace(search, replace interface{}, subject string) string {

	var aSearch, aReplace []string

	aSearch, aReplace, err := buildReplaceSlice(search, replace)
	if err != nil {
		return subject
	}

	r := ""
	for index, s := range aSearch {
		if index < len(aReplace) {
			r = aReplace[index]
		}
		subject = strings.Replace(subject, s, r, -1)
	}

	return subject
}

// Ireplace is case-insensitive version of Replace()
func Ireplace(search, replace interface{}, subject string) string {

	var aSearch, aReplace []string

	aSearch, aReplace, err := buildReplaceSlice(search, replace)
	if err != nil {
		return subject
	}

	r := ""
	var reg *regexp.Regexp
	for index, s := range aSearch {
		if index < len(aReplace) {
			r = aReplace[index]
		}
		reg = regexp.MustCompile("(?i:" + s + ")")
		subject = reg.ReplaceAllString(subject, r)
	}

	return subject
}

// Addslashes quote string with slashes
//
// The characters to be escaped are single quote ('), double quote (") and backslash (\).
func Addslashes(str string) string {
	return Replace([]string{"\\", "'", "\""}, []string{"\\\\", "\\'", "\\\""}, str)
}

// Stripslashes un-quotes a quoted string
func Stripslashes(str string) string {
	return Replace([]string{"\\\\", "\\'", "\\\""}, []string{"\\", "'", "\""}, str)
}

// Chr returns a one-character string containing the character specified by ascii
func Chr(ascii int) string {
	for ascii < 0 {
		ascii += 256
	}
	ascii %= 256
	return string(ascii)
}

// Ord return ASCII value of character
func Ord(character string) rune {
	return []rune(character)[0]
}

// Explode returns an slice of strings, each of which is a substring of str
// formed by splitting it on boundaries formed by the string delimiter.
func Explode(delimiter, str string) []string {
	return strings.Split(str, delimiter)
}

// Implode returns a string containing a string representation of all the slice
// elements in the same order, with the glue string between each element.
func Implode(glue string, pieces []string) string {
	return strings.Join(pieces, glue)
}

// Lcfirst make a string's first character lowercase
func Lcfirst(str string) string {
	return strings.ToLower(Substr(str, 0, 1)) + Substr(str, 1, 0)
}

// Ucfirst make a string's first character uppercase
func Ucfirst(str string) string {
	return strings.ToUpper(Substr(str, 0, 1)) + Substr(str, 1, 0)
}

// Md5 calculate the md5 hash of a string
func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// Md5File calculates the md5 hash of a given file
func Md5File(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()

	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

// Strstr find the first occurrence of a string
func Strstr(haystack, needle string) string {
	pos := Strpos(haystack, needle)
	if pos < 0 {
		return ""
	}
	return Substr(haystack, pos, 0)
}

// Stristr is case-insensitive Strstr()
func Stristr(haystack, needle string) string {
	pos := Stripos(strings.ToLower(haystack), strings.ToLower(needle))
	if pos < 0 {
		return ""
	}
	return Substr(haystack, pos, 0)
}

// Crc32 calculates the crc32 polynomial of a string
func Crc32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

// Bin2hex converts binary data into hexadecimal representation
func Bin2hex(src []byte) string {
	return hex.EncodeToString(src)
}

// Hex2bin decodes a hexadecimally encoded binary string
func Hex2bin(str string) []byte {
	s, _ := hex.DecodeString(str)
	return s
}

// Trim strips whitespace (or other characters) from the beginning and end of a string
func Trim(str string, chars ...string) string {
	if len(chars) == 0 {
		return strings.TrimSpace(str)
	}
	var cutset = ""
	for _, c := range chars {
		cutset += c
	}
	return strings.Trim(str, cutset)
}

// Ltrim strips whitespace (or other characters) from the beginning of a string
func Ltrim(str string, chars ...string) string {
	if len(chars) == 0 {
		return strings.TrimLeftFunc(str, unicode.IsSpace)
	}
	var cutset = ""
	for _, c := range chars {
		cutset += c
	}
	return strings.TrimLeft(str, cutset)
}

// Rtrim strips whitespace (or other characters) from the end of a string
func Rtrim(str string, chars ...string) string {
	if len(chars) == 0 {
		return strings.TrimRightFunc(str, unicode.IsSpace)
	}
	var cutset = ""
	for _, c := range chars {
		cutset += c
	}
	return strings.TrimRight(str, cutset)
}
