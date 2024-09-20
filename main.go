package main

import "fmt"

func customerData( openingBalance, itemsCharged, creditsApplied, creditAllowed int)(int, string) {
	newBalance := openingBalance + itemsCharged - creditsApplied

	var message string
	if newBalance > creditAllowed {
		message = "credit limit exceded"
	} else{
		message = "credit not exceeded"
	}
	return newBalance, message
}

func maino() {
	var accountNumber, openingBalance, itemsCharged, creditApplied, creditAllowed int

	fmt.Print("Account number? ")
	fmt.Scan(&accountNumber)

	fmt.Print("Balance at the begining of the month? ")
	fmt.Scan(&openingBalance)

	fmt.Print("Total of all items charged by the customer this month? ")
	fmt.Scan(&itemsCharged)

	fmt.Print("Total credits applied to the customers account this month? ")
	fmt.Scan(&creditApplied)

	fmt.Print("Allowed credit limit? ")
	fmt.Scan(&creditAllowed)

	 balance, message := customerData(openingBalance,itemsCharged,creditApplied,creditAllowed)

	 fmt.Println("Your balance is ", balance)
	 fmt.Print("Message", message)
}