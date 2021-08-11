package conversion

import "strconv"

func SliceStringToFloat64(values []string) []float64 {
	var newString []float64
	for _, val := range values {
		newVal, _ := strconv.ParseFloat(val, 64)

		newString = append(newString, newVal)
	}
	return newString
}
