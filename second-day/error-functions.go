package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	fPath := "/tmp/dat1"
	inD := []byte("hello\nWorld")
	if err := ioutil.WriteFile(fPath, inD, 0644); err != nil {
		panic(err)
	}
	fmt.Println("Write was successful!")
	outData, err := ioutil.ReadFile(fPath + "x")
	if err != nil {
		panic(err)
	}
	fmt.Print(string(outData))
}
