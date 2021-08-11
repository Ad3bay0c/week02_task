package main

import (
	"reflect"
	"testing"
)

func TestCalculateTable(t *testing.T) {
	tables := []struct {
		values []string
		output []float64
		word   string
	}{
		{
			[]string{"2*3*1", "10/2/5", "10-4-2", "15+12"},
			[]float64{6, 1, 4, 27},
			"First result",
		},
		{
			[]string{"3*3*1", "20/2/5", "-10-4-2", "15+12+10"},
			[]float64{9, 2, -16, 37},
			"Second result",
		},
		{
			[]string{"0*2*3*1", "0*10/2/5", "-10-4-2", "15+12"},
			[]float64{0, 0, -16, 27},
			"Third result",
		},
	}

	for _, table := range tables {
		t.Run(table.word, func(t *testing.T) {
			result := Calculate(table.values...)

			if !reflect.DeepEqual(result, table.output) {
				t.Errorf("Want: %v, Got: %v ", table.output, result)
			}
		})
	}
}