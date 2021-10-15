package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
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

func ToUpper(s string) string {
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

func aToAn(s string) []string {
	splitString := strings.Split(s, " ")
	// this splits the string into its own elements by the space.
	for i, ele := range splitString {
		// the following code is to convert a/A into an/An
		if i >= 0 && i < len(splitString)-1 {
			// since i'm dealing with i+1, i dont want to go outside of the range so i the
			// the last the index the code should consider is the penultimate index hence
			// i<len(splitString-1 which is equivalent to i==len(splitString)-2)
			if strings.HasPrefix(splitString[i+1], "a") || strings.HasPrefix(splitString[i+1], "e") || strings.HasPrefix(splitString[i+1], "i") || strings.HasPrefix(splitString[i+1], "o") || strings.HasPrefix(splitString[i+1], "u") || strings.HasPrefix(splitString[i+1], "h") {
				// strings.HasPrefix checks if the string starts with the substring
				// the ele in splitString[i] is the string and "a" is the substring.
				// then repeat for all lowercase vowels and h.
				if ele == "a" {
					splitString[i] = "an"
				}
				if ele == "A" {
					splitString[i] = "An"
				}
			}
			if strings.HasPrefix(splitString[i+1], "A") || strings.HasPrefix(splitString[i+1], "E") || strings.HasPrefix(splitString[i+1], "I") || strings.HasPrefix(splitString[i+1], "O") || strings.HasPrefix(splitString[i+1], "U") || strings.HasPrefix(splitString[i+1], "H") {
				// strings.HasPrefix checks if the string starts with the substring
				// the ele in splitString[i] is the string and "a" is the substring.
				// then repeat for all uppercase vowels and H.
				if ele == "a" {
					splitString[i] = "an"
				}
				if ele == "A" {
					splitString[i] = "An"
				}
			}
			// the above code was to change a/A to an/An.
		}
	}
	return splitString
}

func Hex(s string) string {
	if num, err := strconv.ParseInt(s, 16, 64); err == nil {
		// this converts the string into an int64
		// int64 means that there is a binary code with 64 places, meaning it
		// accounts for a greater number of characters/numbers whatever.
		// we use base 16 for hexdecimal because each hexdecimal value is
		// made up of 16 inputs:
		// 0,1,2,3,4,5,6,7,8,9,a,b,c,e,d,f.
		num2 := int(num)
		// we then had to convert the product of num which is an int64 into an int
		itos := strconv.Itoa(num2)
		// then we converted that int back into a string
		s = itos
	}
	return s
}

func bin(s string) string {
	if num, err := strconv.ParseInt(s, 2, 64); err == nil {
		// we use base 2 for bin (binary) because binary only works with 2 inputs 0 and 1
		num2 := int(num)
		itos := strconv.Itoa(num2)
		s = itos
	}
	return s
}

func main() {
	fileName := "sample.txt"
	content, err := ioutil.ReadFile(fileName)
	// this reads the content of the file
	if err != nil {
		fmt.Printf("the mistake is : %v\n", err.Error())
	}
	sampleString := string(content)
	// this turns the contents of the file into a string.
	sliceOfString := aToAn(sampleString)
	// this has converted the a/A's to an/Ans and has turn the sample string into a slice of string
	fmt.Println(sliceOfString)
	for i := range sliceOfString {
		if i >= 0 && i < len(sliceOfString)-1 {
			// since i'm dealing with i+1, i dont want to go outside of the range so i the
			// the last the index the code should consider is the penultimate index hence
			// i<len(splitString-1 which is equivalent to i==len(splitString)-2)
			if strings.HasPrefix(sliceOfString[i+1], "(cap") {
				sliceOfString[i] = Capitalize(sliceOfString[i])
			} // this deals with capitalise, without recursions.

			if strings.HasPrefix(sliceOfString[i+1], "(up") {
				sliceOfString[i] = ToUpper(sliceOfString[i])
			} // this deals with uppercasing the word, without recursions.

			if strings.HasPrefix(sliceOfString[i+1], "(low") {
				sliceOfString[i] = ToLower(sliceOfString[i])
			} // this deals with lowercasing the word, without recursions.

			if strings.HasPrefix(sliceOfString[i+1], "(hex)") {
				sliceOfString[i] = Hex(sliceOfString[i])
			} // this deals with hexadecimal version to the decimal version of the word, without recursions.

			if strings.HasPrefix(sliceOfString[i+1], "(bin)") {
				sliceOfString[i] = bin(sliceOfString[i])
			} // this deals with convert binary version  to decimal numbers, without recursions.

			// for i := range sliceOfString {
			// 	if strings.HasPrefix(sliceOfString[i], "(cap") {
			// 		sliceOfString[i] = strconv.Atoi(sliceOfString[i])
			// 		sliceOfString[i] = strings.Trim(sliceOfString[i], sliceOfString[i])
			// 	}
			// }
			// var newString string
			// for a :=range splitString{
			// 	newString= append(newString,splitString[a])
			// }
			fmt.Println(sliceOfString)
		}
	}
}
