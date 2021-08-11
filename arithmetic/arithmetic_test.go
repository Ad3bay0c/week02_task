package arithmetic

import "testing"

func TestAddTable(t *testing.T) {
	tables := []struct {
		values []float64
		output float64
		word   string
	}{
		{
			[]float64{-2, 0, 4, 6},
			8,
			"Positive Result",
		},
		{
			[]float64{2, 0, 4, 6, 10},
			22,
			"Second Positive Result",
		},
		{
			[]float64{0, 0, 12, -12, -8},
			-8,
			"Negative Result",
		},
	}

	for _, table := range tables {
		t.Run(table.word, func(t *testing.T) {
			add := Add(table.values)

			if add != table.output {
				t.Errorf("Want: %v, Got: %v ", table.output, add)
			}
		})
	}
}

func TestSubTable(t *testing.T) {
	tables := []struct {
		values []float64
		output float64
		word   string
	}{
		{
			[]float64{24, -2, 0, 4, 6},
			16,
			"Positive Result",
		},
		{
			[]float64{12, -10, 22},
			0,
			"Zero Result",
		},
		{
			[]float64{0, 0, -12, 12, 8},
			-8,
			"Negative Result",
		},
	}

	for _, table := range tables {
		t.Run(table.word, func(t *testing.T) {
			sub := Sub(table.values)

			if sub != table.output {
				t.Errorf("Want: %v, Got: %v ", table.output, sub)
			}
		})
	}
}

func TestMultTable(t *testing.T) {
	tables := []struct {
		values []float64
		output float64
		word   string
	}{
		{
			[]float64{4, 2, 1, 6},
			48,
			"Positive Result",
		},
		{
			[]float64{12, 0, -22},
			0,
			"Zero Result",
		},
		{
			[]float64{-12, 12, 8},
			-1152,
			"Negative Result",
		},
	}

	for _, table := range tables {
		t.Run(table.word, func(t *testing.T) {
			mult := Mult(table.values)

			if mult != table.output {
				t.Errorf("Want: %v, Got: %v ", table.output, mult)
			}
		})
	}
}

func TestDivTable(t *testing.T) {
	tables := []struct {
		values []float64
		output float64
		word   string
	}{
		{
			[]float64{24, 2, 6},
			2,
			"Positive Result",
		},
		{
			[]float64{12, -1},
			-12,
			"Zero Result",
		},
		{
			[]float64{-12, -1},
			12,
			"Negative Result",
		},
	}

	for _, table := range tables {
		t.Run(table.word, func(t *testing.T) {
			div := Div(table.values)

			if div != table.output {
				t.Errorf("Want: %v, Got: %v ", table.output, div)
			}
		})
	}
}
