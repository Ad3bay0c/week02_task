package arithmetic

func Sub(operation []float64) float64 {
	sub := float64(0)

	for idx, val := range operation {
		if idx < 1 {
			sub = val - sub
		} else {
			sub -= val
		}
		
	}

	return sub
}