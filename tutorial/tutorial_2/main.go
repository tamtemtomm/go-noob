package main

import (
	"errors"
	"fmt"
)

func main() {
	printMe("Hello World")

	var numerator int = 12
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)

	// if err!= nil{
	// 	fmt.Println(err.Error())
	// } else if remainder == 0{
	// 	fmt.Printf("The result of the integer division is %v", result)
	// } else {
	// 	fmt.Printf("The result of the ineger division is %v with remainder %v", result, remainder)
	// }

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v\n", result)
	default:
		fmt.Printf("The result of the ineger division is %v with remainder %v\n", result, remainder)

	}

	switch remainder {
	case 0:
		fmt.Println("The division was exact")
	case 1, 2: 
		fmt.Println("The division was close")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divided by zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err
}
