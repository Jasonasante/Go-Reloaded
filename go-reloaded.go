package main

import (
	"fmt"
	"io/ioutil"
)

func low() {
}

func up() {
}

func cap() {
}

func atoAn() {
}

func hex() {
}

func main() {
	fileName := "sample.txt"
	content, err := ioutil.ReadFile(fileName)
	// this reads the content of the file
	if err != nil {
		fmt.Printf("the mistake is : %v\n", err.Error())
	}
	sampleString := string(content)
	// this converts the content of the file into a string
	fmt.Println(sampleString)
}
