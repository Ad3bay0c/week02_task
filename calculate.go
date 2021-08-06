package main

import (
	"fmt"
	"strings"

	a "github.com/Ad3bay0c/arithmetic"
	"github.com/Ad3bay0c/conversion"
)
func main() {
	fmt.Println(Calculate("2*-5*2", "10/2/0", "10+3+2+1+3", "14-10-6-3-2-3"))
	fmt.Println(a.Div([]float64{30,2,5}))
	fmt.Println(a.Add([]float64{30,-2,5}))
	fmt.Println(a.Sub([]float64{30,2,-5}))
	fmt.Println(a.Mult([]float64{30,2,5}))

	fmt.Println(Calculate([]string{"3*4*4", "4+2+1"}...))
}

func Calculate(expression ...string) []float64 {
	var result []float64
	for _, val := range expression {
		// Check for each operator and call corresponding Method
		switch {
		case strings.Contains(val, "*"):
			s := strings.Split(val, "*") // []string
			newVal := conversion.SliceStringToFloat64(s) // Converts the slice of strings to slice of floats
			result = append(result, a.Mult(newVal))

		case strings.Contains(val, "+"):
			s := strings.Split(val, "+") // []string
			newVal := conversion.SliceStringToFloat64(s)
			result = append(result, a.Add(newVal))

		case strings.Contains(val, "/"):
			s := strings.Split(val, "/") // []string
			newVal := conversion.SliceStringToFloat64(s)
			result = append(result, a.Div(newVal)) 

		case strings.Contains(val, "-"):
			s := strings.Split(val, "-") // []string
			newVal := conversion.SliceStringToFloat64(s)
			result = append(result, a.Sub(newVal))
		}
	}

	return result
}