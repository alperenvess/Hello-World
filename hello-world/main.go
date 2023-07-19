/// String in Golang

package main

import (
	"fmt"
)

func main() {

	// 1 Works
	// var myName string = "myName"
	// var myName = "Alperen"

	// 2 Works
	// myName := "Alperen"
	// fmt.Println("Hello!", aurora.Magenta(myName))

	// 3 Works
	// var myName
	// myName = "Alperen"

	//4 Works
	// var myName string
	// myName = "Alperen"

	//5  Dosen't Work
	// var myName string
	// myName := "Alperen" (Because we already defined the variable on the upper line we dont need to use the : for defining)

	// var Country, City string
	// Country = "Turkey"
	// City = "Konya"

	// fmt.Println(Country, City)

	HelloWorld := []string{"dzień dobry", "Hallo", "Guten Tag"}
	fmt.Printf("A: len: %d cap: %d \n", len(HelloWorld),
		cap(HelloWorld))

	Czechia := "Ahoy"
	HelloWorld = append(HelloWorld, Czechia)

	fmt.Printf("B: len: %d cap: %d \n", len(HelloWorld),
		cap(HelloWorld))

	fmt.Printf("%v\n", HelloWorld)
	fmt.Printf("%v\n", HelloWorld)

}
