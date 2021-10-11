package piscine

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

func readFile() string {
	fileName := "sample.txt"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("the mistake is : %v\n", err.Error())
	}
	var bytes byte
	for i := range content {
		bytes += content[i]
	}
	return string(bytes)
}
