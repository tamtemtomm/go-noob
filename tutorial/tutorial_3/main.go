package main

import "fmt"

// ARRAY, SLICE, AND LOOPS

func main() {
	// ARRAY
	var intArr [3]int32
	intArr[1] = 1
	fmt.Println(intArr)
	fmt.Println(intArr[1:])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intArr2 [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr2)

	intArr3 := [...]int32{1, 2, 3}
	fmt.Println(intArr3)

	// SLICES
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity of %v", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v with capacity of %v", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("Created new array with %v length and %v capacity\n", len(intSlice3), len(intSlice3))

	// MAP (KEY : VALUE {object or dictionary})
	var myMap map[string]uint16 = make(map[string]uint16)
	fmt.Println(myMap)

	var myMap2 = map[string]uint16{"Adam": 23, "Sarah": 45, "Timo":20}
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Jason"])

	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Printf("No age found\n")
	}

	delete(myMap2, "Adam")
	fmt.Println(myMap2)

	// LOOP
	for name:= range myMap2{
		fmt.Printf("Name: %v\n", name)
	}

	for name, age:= range myMap2{
		fmt.Printf("Name: %v | Age : %v\n", name, age)
	}

	var i int16 = 0
	for i < 10{
		fmt.Println(i)	
		i = i + 1
	}

	for i:=0; i<10; i++{
		fmt.Println(i)
	}
}
