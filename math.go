package php

import (
	"fmt"
	"math"
	"strconv"
)

// Round returns the rounded value of val to specified precision (number of digits after the decimal point)
//
// Note that we limit the range of precision to [-14, 14]
func Round(val float64, precision int) float64 {

	if val == 0 {
		return 0
	}

	pLimit := 14

	if precision > pLimit {
		precision = pLimit
	} else if precision < -pLimit {
		precision = -pLimit
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
