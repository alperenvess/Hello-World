package main

import (
	"errors"
	"fmt"
)

var errUnknownOperation = errors.New("Unknown operation")

func Calculate(op string, a int, b int) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	}
	return 0, errUnknownOperation
}

func main() {
	r, _ := Calculate("+", 1, 2)
	fmt.Println(r)
}
