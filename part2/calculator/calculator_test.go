package calculator

import (
	"reflect"
	"testing"
)

func TestCalculatorBracketSumSubBehaviour(t *testing.T) {
	input := "(1+2)-3\n"

	var rExpected float64 = 0

	r, e := Calculate(input)

	if e != nil {
		t.Errorf("Error was not expected")
	}

	if r != rExpected {
		t.Errorf("Check failed. Got %f\nExpected %f", r, rExpected)
	}
}

func TestCalculatorBracketSumSubDivMultBehaviour(t *testing.T) {
	input := "(2*10+1)/(4+6/2)\n"

	var rExpected float64 = 3

	r, e := Calculate(input)

	if e != nil {
		t.Errorf("Error was not expected")
	}

	if r != rExpected {
		t.Errorf("Check failed. Got %f\nExpected %f", r, rExpected)
	}
}

func TestCalculatorWhiteSpaceBehaviour(t *testing.T) {
	input := "( 2 * 10 + 1 ) / ( 4 + 6 / 2 )\n"

	var rExpected float64 = 3

	r, e := Calculate(input)

	if e != nil {
		t.Errorf("Error was not expected")
	}

	if r != rExpected {
		t.Errorf("Check failed. Got %f\nExpected %f", r, rExpected)
	}
}

func TestPolishNotationSumSubBehaviour(t *testing.T) {
	input := "(2*10+1)/(4+6/2)\n"

	pnExpected := []string{"2", "10", "*", "1", "+", "4", "6", "2", "/", "+", "/"}

	pn, e := getPolishNotation(input)

	if e != nil {
		t.Errorf("Error was not expected")
	}

	if !reflect.DeepEqual(pn, pnExpected) {
		t.Errorf("Check failed. Got %s\nExpected %s", pn, pnExpected)
	}
}

func TestPolishNotationBracketSumSubDivMulBehaviour(t *testing.T) {
	input := "1+2-1\n"

	pnExpected := []string{"1", "2", "+", "1", "-"}

	pn, e := getPolishNotation(input)

	if e != nil {
		t.Errorf("Error was not expected")
	}

	if !reflect.DeepEqual(pn, pnExpected) {
		t.Errorf("Check failed. Got %s\nExpected %s", pn, pnExpected)
	}
}
