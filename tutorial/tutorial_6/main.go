package main

import "fmt"

// POINTER

func main() {

	// POINTER
	var p *int32 = new(int32)
	var i int32

	fmt.Printf("The value of p points to is %v\n", *p)
	fmt.Printf("The value if i is : %v\n", i)

	p = &i
	*p = 1

	fmt.Printf("The value of p points to is %v\n", *p)
	fmt.Printf("The value if i is : %v\n", i)

	var k int32 = 2
	i = k

	// SLICES POINTER
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	// ARRAY FUNCTION (THEY CREATE A COPY)
	var thing = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing1 array is %p\n", &thing)
	var result [5]float64 = square(thing)
	fmt.Printf("The result is: %v\n", result)

	// ARRAY FUNCTION (THEY CREATE A COPY)
	fmt.Printf("\nThe memory location of the thing2 array is %p\n", &thing)
	var result2 [5]float64 = squarePointer(&thing)
	fmt.Printf("The result is: %v\n", result2)
}

func square(thing1 [5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing1 array is %p\n", &thing1)
	for i := range thing1 {
		thing1[i] = thing1[i] * thing1[i]
	}
	return thing1
}

func squarePointer(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is %p\n", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}

