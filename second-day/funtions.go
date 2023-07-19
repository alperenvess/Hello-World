package main

import (
	"fmt"
	"strings"
)

func printThings(letter []string, decorate func(string) string) {
	for _, l := range letter {
		d := decorate(l)
		fmt.Println(d)
	}
}

func main() {
	alphabet := []string{"mleko", "cars", "programming"}

	printThings(alphabet, func(a string) string {
		return strings.ToUpper(a)

	})

}
