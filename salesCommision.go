package main

import "fmt"

func calculate(itemSold string, numberOfSoldItems int)(int){
	var comissionRate float64 = 9.0 / 100.0
	var salary int = 200

	var item1 float64 = 239.99
	var item2 float64 = 129.75
	var item3 float64 = 99.95
	var item4 float64 = 350.89

	var commision float64 = 0.0

	if itemSold == "1" {
		commision =  comissionRate * (item1 * float64(numberOfSoldItems))
	}
	if itemSold == "2" {
		commision =  comissionRate * (item2 * float64(numberOfSoldItems))
	}
	if itemSold == "3" {
		commision =  comissionRate * (item3 * float64(numberOfSoldItems))
	}
	if itemSold == "4" {
		commision =  comissionRate * (item4 * float64(numberOfSoldItems))
	}else{
		fmt.Println("invalid item")
		return salary
	}
	return salary + int(commision)
}
