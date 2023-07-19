package main

import (
	"fmt"
	"testing"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(10, 20))
}

func TestAdd(t *testing.T) {
	if add(10, 25) != 20 {
		t.Fatal("Boom")
	}
}
