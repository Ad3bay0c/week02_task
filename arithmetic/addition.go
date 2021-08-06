package arithmetic

func Add(operation []float64) float64 {
		if len(operation) == 0 {
			return 0
		}
		// using Recursion
	return Add(operation[1:])+ operation[0]
}
