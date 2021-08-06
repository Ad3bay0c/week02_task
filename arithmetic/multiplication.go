package arithmetic

func Mult(operation []float64) float64 {
	// mult := float64(1)

	// for _, val := range operation {
	// 	mult *= val
	// }
	if len(operation) == 0 {
		return 1
	}

	return operation[0] * Mult(operation[1:])
}
