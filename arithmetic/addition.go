package arithmetic

func Add(operation []float64) float64 {
	sum := float64(0)

	for _, val := range operation {
		sum += val
	}

	return sum
}