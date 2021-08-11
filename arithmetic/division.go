package arithmetic

import "fmt"

func Div(operations []float64) float64 {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("Oops!!!, This happened: %v\n", err)
		}
	}()

	div := float64(1)

	for idx, val := range operations {
		
		if idx < 1 {
			div = val / div
		} else if val == 0 {
			panic("Division By Zero Not Allowed")
		} else {
			div /= val
		}
		
	}

	return div
	
}