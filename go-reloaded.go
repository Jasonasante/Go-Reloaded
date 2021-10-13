package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ToLower(s string) string {
	lowstr := []rune(s)
	for i, char := range lowstr {
		if char >= 65 && char <= 90 {
			lowstr[i] = lowstr[i] + 32
		}
	}
	return string(lowstr)
}

func ToUpper(s string) {
	upstr := []rune(s)
	for i, char := range upstr {
		if char >= 97 && char <= 122 {
			upstr[i] = upstr[i] - 32
		}
	}
	return string(upstr)
}

func Capitalize(s string) string {
	var cap []rune
	l := ToLower(s)
	// lowers the incoming string
	r := []rune(l)
	// the now lowercase string is turn into a slice of runes
	if len(l) == 1 {
		if r[0] >= 'a' && r[0] <= 'z' {
			r[0] = r[0] - 32
		}
		// this is for one letter.
	}
	for i := 0; i < len(l)-1; i++ {
		if r[0] >= 'a' && r[0] <= 'z' {
			r[0] = r[0] - 32
		}
		// this is for the first letter of the word
		if r[i] < 48 || r[i] > 57 && r[i] < 65 || r[i] > 90 && r[i] < 97 || r[i] > 122 {
			if r[i+1] >= 'a' && r[i+1] <= 'z' {
				r[i+1] = r[i+1] - 32
			}
			// this take into account if there are multiple words and will capitalise each
			// word after a non-alpha-numerical character.
		}
		cap = append(cap, r[i])
		// this appends every each character at each index into cap (empty slice of runes).
		// So characters after the first letter of each word are not capitalised and are
		// appended into cap.
	}
	cap = append(cap, r[(len(r)-1)])
	// this is to append the last index at the end, because the previous line of code looks
	// at r[i+1], and since there is no r[i+1] at the end the code will produce an error.
	return string(cap)
}

// func aToAn(s string)string {
//  strings.Replace()
// }

// func hex() {
// }

func main() {
	fileName := "sample.txt"
	content, err := ioutil.ReadFile(fileName)
	// this reads the content of the file
	if err != nil {
		fmt.Printf("the mistake is : %v\n", err.Error())
	}
	sampleString := string(content)
	// this turns the contents of the file into a string.
	splitString := strings.Split(sampleString, " ")
	// this splits the string into its own elements by the space.
	for i := range splitString {
		if splitString[i] == "(cap)" {
		}
	}
	// this converts the content of the file into a string
	// toFindCap:= "(cap)"
	// sampleSlice:=[]rune(sampleString)
	// subCap:=[]rune(toFindCap)
	// counter:=0
	// for i:=0; i<len(sampleString)-len(toFindCap); i++{
	// 	for j:=0 ; j<len(toFindCap); j++{
	// 		if

	// 	}
	// }
}
