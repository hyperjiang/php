package php

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	"golang.org/x/exp/constraints"
)

// Real is a type constraint for real number
type Real interface {
	constraints.Integer | constraints.Float
}

// MaxPrecision is the default max precision for math functions
var MaxPrecision = 14

// Round returns the rounded value of val to specified precision (number of digits after the decimal point)
//
// Note that we limit the range of precision to [-14, 14] by default
func Round(val float64, precision int) float64 {

	if val == 0 {
		return 0
	}

	if precision > MaxPrecision {
		precision = MaxPrecision
	} else if precision < -MaxPrecision {
		precision = -MaxPrecision
	}

	factor := math.Pow10(int(math.Abs(float64(precision))))
	if precision < 0 {
		return Round(val/factor, 0) * factor
	}

	format := fmt.Sprintf("%%.%df", precision)
	if val > 0 {
		val += 0.05 / factor
	} else {
		val -= 0.05 / factor
	}
	v := fmt.Sprintf(format, val)
	res, _ := strconv.ParseFloat(v, 64)
	return res
}

// Abs returns the absolute value of number
func Abs[T Real](number T) T {
	return T(math.Abs(float64(number)))
}

// MtRand is alias of Rand()
func MtRand(min, max int) int {
	return Rand(min, max)
}

// Rand generates a random integer
func Rand(min, max int) int {
	if min > max {
		return 0
	}
	m := int(Abs(float64(min)))
	return rand.Intn(max+m) - m
}

// BaseConvert converts a number between arbitrary bases
func BaseConvert(number string, frombase, tobase int) (string, error) {
	i, err := strconv.ParseInt(number, frombase, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, tobase), nil
}

// Bindec converts a binary number to an integer
func Bindec(str string) (string, error) {
	return BaseConvert(str, 2, 10)
}

// Decbin converts decimal to binary
func Decbin(number int64) string {
	return strconv.FormatInt(number, 2)
}

// Dechex converts decimal to hexadecimal
func Dechex(number int64) string {
	return strconv.FormatInt(number, 16)
}

// Hexdec converts hexadecimal to decimal
func Hexdec(str string) (int64, error) {
	return strconv.ParseInt(str, 16, 0)
}

// Decoct converts decimal to octal
func Decoct(number int64) string {
	return strconv.FormatInt(number, 8)
}

// Octdec converts octal to decimal
func Octdec(str string) (int64, error) {
	return strconv.ParseInt(str, 8, 0)
}

// Ceil returns the next highest integer value by rounding up value if necessary
func Ceil(value float64) float64 {
	return math.Ceil(value)
}

// Floor returns the next lowest integer value by rounding down value if necessary
func Floor(value float64) float64 {
	return math.Floor(value)
}

// Pi gets value of pi
func Pi() float64 {
	return math.Pi
}
