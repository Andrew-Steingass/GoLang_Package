// Package trimmean provides functions to calculate trimmed means for slices of floats.
package trimmean

import (
	"math"
	"sort"
)

// TrimmedMean calculates the trimmed mean of a float64 slice.
// If one trim argument is provided, trimming is symmetric.
// If two trim arguments are provided, the first is for lower trimming, the second for upper.
func TrimmedMean(slice []float64, trimArgs ...float64) float64 {
	var lowerTrim, upperTrim float64
	if len(trimArgs) == 1 {
		lowerTrim, upperTrim = trimArgs[0], trimArgs[0]
	} else if len(trimArgs) == 2 {
		lowerTrim, upperTrim = trimArgs[0], trimArgs[1]
	} else {
		panic("Provide 1 or 2 trimming arguments")
	}

	sort.Float64s(slice)
	n := len(slice)
	lowerCount := int(lowerTrim * float64(n))
	upperCount := int(upperTrim * float64(n))
	trimmedSlice := slice[lowerCount : n-upperCount]

	var sum float64
	for _, value := range trimmedSlice {
		sum += value
	}

	return sum / float64(len(trimmedSlice))
}

// Round rounds a float64 value to the specified number of decimal places.
func Round(value float64, decimals int) float64 {
	multiplier := math.Pow(10, float64(decimals))
	return math.Round(value*multiplier) / multiplier
}
