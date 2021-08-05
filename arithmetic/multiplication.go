package arithmetic

func Mult(operation []float64) float64 {
	mult := float64(1)

	for _, val := range operation {
		mult *= val
	}

	return mult
}
