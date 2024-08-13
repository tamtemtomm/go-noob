package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 5

func main(){
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites{
		go checkChickenPrice(websites[i], chickenChannel)
		go checkTofuPrice(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrice(website string, chickenChannel chan string){
	for{
		time.Sleep(time.Second*1)
		var chickenPrice = rand.Float32()*20
		if chickenPrice <= MAX_CHICKEN_PRICE{
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrice(website string, tofuChannel chan string){
	for{
		time.Sleep(time.Second*1)
		var tofuPrice = rand.Float32()*20
		if tofuPrice <= MAX_TOFU_PRICE{
			tofuChannel <- website
			break
		}
	}
}


func sendMessage(chickenChannel chan string, tofuChannel chan string){
	select{
	case website := <-chickenChannel:
		fmt.Printf("\nFound a deal on chicken at %v", website)
	case website := <-tofuChannel:
		fmt.Printf("\nEmail send : Found a deal on tofu at %v", website)
	}
}