package main

import (
	"fmt"
	"os"
)

// func nameByKey(names map[string]string, key string) *string {

// 	var result *string

// 	for k, v := range names {
// 		if key == k {
// 			/// NEVER DO IT
// 			result = &v
// 		}
// 	}
// 	return result
// }
/////////// Mapping //////
// func main() {
// 	names := map[string]string{
// 		"12": "mike",
// 		"34": "joe",
// 		"99": "Joanna",
// 	}

// 	found := nameByKey(names, "12")
// 	fmt.Printf("&s", *found)
// }

// func main() {

// 	var milk string = "우유"

// 	for index, runeValue := range milk {
// 		fmt.Printf("%c (%U) starts at byte position %d\n",
// 			runeValue, runeValue, index)
// 	}

// 	for i := 0; i < len(milk); i++ {
// 		fmt.Printf("byte: %x at the index %d\n",
// 			milk[i], i)
// 	}
// }

// func main() {

// 	// helloMsgs := map[string]string{}
// 	// helloMsgs["pl"] = "Dzień Dobry"
// 	// helloMsgs["en"] = "Good Morning"
// 	// helloMsgs["tr"] = "Günaydın"

// 	// fmt.Printf("%v", helloMsgs["tr"])

// 	type daytimeGreetings map[string][]string
// 	type g map[string]daytimeGreetings

// 	helloMsgs := map[string]daytimeGreetings{
// 		"pl": {"morning": []string{
// 			"Dzień Dobry",
// 		}},
// 		"en": {"morning": []string{
// 			"Good Morning",
// 		}},
// 		"tr": {"morning": []string{
// 			"Günaydın",
// 		}},
// 	}

// 	var lang string
// 	fmt.Scan(&lang)

// 	if val, ok := helloMsgs[lang]; ok {
// 		fmt.Printf("%v", val)

// 	} else {

// 	}

// }

// type myInt int

// func display(i int) {
// 	fmt.Printf("%d", i)
// }

// func main() {
// 	var i myInt = 12
// 	i = i + 12
// 	// to fix it we should declare the type of the "i"
// 	display(int(i))
// }

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("struct1: %v\n", p)
	fmt.Printf("struct2: %+v\n", p)
	fmt.Printf("struct3: %#v\n", p)
	fmt.Printf("type: %T\n", p)
	fmt.Printf("bool: %t\n", true)
	fmt.Printf("int: %d\n", 123)
	fmt.Printf("bin: %b\n", 14)
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
