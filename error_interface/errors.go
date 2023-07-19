package main

import "fmt"

func main() {


func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Can not divide '%d' by zero", a)
	}
	return a / b, nil
}
}