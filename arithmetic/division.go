package arithmetic

func Div(operation []float64) float64 {

	div := float64(1)

	for idx, val := range operation {
		if idx < 1 {
			div = val / div
		} else {
			div /= val
		}
		
	}

	return div
	
}