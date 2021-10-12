package main

import (
	"fmt"
	"io/ioutil"
)

func low(s string, num int) {
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
	toFindCap:= "(cap)"
	sampleSlice:=[]rune(sampleString)
	subCap:=[]rune(toFindCap)
	counter:=0
	for i:=0; i<len(sampleString)-len(toFind); i++{
		for j:=0 ; j<len(toFindCap); 
	}
}
