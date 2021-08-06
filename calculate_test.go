package main 

import (
	"testing"
	a "github.com/Ad3bay0c/arithmetic"
)


func TestCalculate(t *testing.T){
	result := Calculate("2*3*4*6*10", "4/2/1/4/5", "4+1+0+1+3", "10-23")

	if len(result) != 4 {
		t.Errorf("Result not accurate")
	}
}

func TestAdd(t *testing.T) {
	add := a.Add([]float64{-2,0,4,6})

	if add != 8 {
		t.Errorf("Sum of Integers not correct")
	}
}

func TestSub(t *testing.T) {
	add := a.Sub([]float64{-2,0,4,4})

	if add != -10 {
		t.Errorf("Inaccurate Calculation")
	} 
}

func TestMult(t *testing.T) {
	add := a.Mult([]float64{-2,0,4,4})

	if add != 0 {
		t.Errorf("Inaccurate Calculation")
	} 
}

func TestDiv(t *testing.T) {
	add := a.Div([]float64{4,-4})

	if add != -1 {
		t.Errorf("Inaccurate Calculation")
	} 
}