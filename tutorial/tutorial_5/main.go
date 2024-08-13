package main

import (
	"fmt"
)

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8{
	return e.gallons*e.mpg
}

func (e electricEngine) milesLeft() uint8{
	return e.mpkwh*e.kwh
}

type engine interface{
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8){
	if miles<=e.milesLeft(){
		fmt.Println("You can make it there")
	} else {
		fmt.Println("Need to fuel up first")
	}
}

func main() {
	var myEngine gasEngine
	fmt.Println(myEngine)

	var myEngine2 gasEngine = gasEngine{mpg: 25, gallons: 15}
	myEngine2.mpg = 20
	fmt.Println(myEngine2)

	var myEngine3 gasEngine = gasEngine{15, 10, owner{"Alex"}}
	fmt.Println(myEngine3.mpg, myEngine3.gallons, myEngine3.owner.name)

	var myEngine4 = struct {
		mpg     uint8
		gallons uint8
	}{15, 5}
	fmt.Println(myEngine4.mpg)

	fmt.Printf("Total miles left in tank: %v\n", myEngine3.milesLeft())

	canMakeIt(myEngine2, 65)

	var myEngine5 electricEngine = electricEngine{20, 100}
	fmt.Println(myEngine5.milesLeft())

	canMakeIt(myEngine5, 100)

}
