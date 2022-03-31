package main

import (
	"fmt"
	"math"
)

func CurrencyCalculator(oneDollar, fiftyCent, twentyCent, tenCent, fiveCent float64) (float64, float64, float64) {
	//Insert your code here
	totalAmount := (oneDollar * 1) + (fiftyCent * 0.5) + (twentyCent * 0.2) + (tenCent * 0.1) + (fiveCent * 0.05)
	twoDollarNotes := math.Floor(totalAmount / 2)
	changeAmount := math.Round((totalAmount-twoDollarNotes*2)*100) / 100
	//Do not remove this
	fmt.Println("Total:", totalAmount, "Two Dollar Notes:", twoDollarNotes, " Change: ", changeAmount)
	return totalAmount, twoDollarNotes, changeAmount
}

func main() {
	var oneDollar float64
	var fiftyCent float64
	var twentyCent float64
	var tenCent float64
	var fiveCent float64

	fmt.Println("Enter the number of 1-dollar coins: ")
	fmt.Scanln(&oneDollar)
	fmt.Println("Enter the number of 50-cent coins: ")
	fmt.Scanln(&fiftyCent)
	fmt.Println("Enter the number of 20-cent coins: ")
	fmt.Scanln(&twentyCent)
	fmt.Println("Enter the number of 10-cent coins: ")
	fmt.Scanln(&tenCent)
	fmt.Println("Enter the number of 5-cent coins: ")
	fmt.Scanln(&fiveCent)
	CurrencyCalculator(oneDollar, fiftyCent, twentyCent, tenCent, fiveCent)
}
