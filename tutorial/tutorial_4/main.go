package main

import (
	"fmt"
	"strings"
)

// STRING, RUNES, AND BYTES
func main() {
	var myString = []rune("résumé")
	fmt.Println(myString)

	var indexed = myString[0]
	fmt.Printf("%v, %T", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	var myRune = 'a'
	fmt.Printf("\nmyRune : %v", myRune)

	var strSlice = []string{"h", "e", "l", "l", "0"}
	var strBuilder strings.Builder
	for i:= range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("\n%v", catStr)
}
