package main

import (
	"fmt"
)

func main() {
	// helloWorld := []string{"dzień dobry", "Ahoj", "Goodmorning"}

	// centralEuHello := helloWorld[0:2]

	// fmt.Printf("central EU len: %d cap: %d \n", len(centralEuHello),
	// 	cap(centralEuHello))
	// we put in the same array
	// centralEuHello = append(centralEuHello, "Good evening 1")

	// we need new one
	// centralEuHello = append(centralEuHello, "Good evening 2")
	// fmt.Printf("central EU len: %d cap: %d \n", len(centralEuHello),
	// 	cap(centralEuHello))

	// fmt.Printf("All: %v\n", helloWorld)
	// fmt.Printf("Central: %v\n", centralEuHello)

	const adress = "ul. Przeskok 2"

	fmt.Printf("len: %d\n", len(adress))

	fmt.Printf("1: %s\n", adress[0:3])
	fmt.Printf("2: %c\n", adress[2])
	fmt.Printf("3: %s\n", adress[2:])
	fmt.Printf("4: %s\n", adress[5:])
	// fmt.Printf("5: %s\n", adress[16:])
	// fmt.Printf("6: %s\n", adress[:16])

}
