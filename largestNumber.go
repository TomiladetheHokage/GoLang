package main

import "fmt"

func largest(numbers... int) int {
	var max int = numbers[0]
	for i:= 0; i < len(numbers); i++{
		if numbers[i] > max{
			max = numbers[i]
		}
	}

	return max
}

func main() {

	var numbers[10]int 
	var number int 

	for i:= 0; i < 5; i++{
		fmt.Println("Enter a number: ")
		fmt.Scan(&number)
		numbers[i] = number
	}
	
	// for  num := range numbers{
	// 	fmt.Printf("%d", num)
	// }
	fmt.Print(largest(number))
	
}