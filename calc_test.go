package main

import "testing"

func TestAdd(t *testing.T) {
	exp := 5
	res := addition(2, 3)
	if res != exp {
		t.Errorf("%d was expect but got %d.\n", exp, res)
	}
}

func TestSubtract(t *testing.T) {
	exp := 2
	res := subtract(7, 5)
	if res != exp {
		t.Errorf("%d was expected but got %d.\n", exp, res)
	}
}

func TestMultiply(t *testing.T) {
	exp := 6
	res := multiply(2, 3)
	if res != exp {
		t.Errorf("%d was expected but got %d.\n", exp, res)
	}
}

func TestDivision(t *testing.T) {
	exp := 2
	res := division(8, 4)
	if res != exp {
		t.Errorf("%d was expected but got %d.\n", exp, res)
	}
}
