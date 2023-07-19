package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	testCases := map[string]struct {
		op string
		a  int
		b  int

		expected int
	}{
		"simple add": {"+", 1, 3, 4},
	}
	for name, v := range testCases {
		t.Logf("test: %s", name)
		r, err := Calculate(v.op, v.a, v.b)
		if err != nil {
			t.Fatalf("%v", err)
		}
		if r != v.expected {
			t.Fatalf("Failed!")
		}
	}
}
