/*
Task:
	Write a program which prompts the user to enter a floating point number 
	and prints the integer which is a truncated version of the floating point
	number that was entered. Truncation is the process of removing the digits to
	the right of the decimal place.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var userInput string

	fmt.Printf("Type a floating point number please:\n")
	_, scanError := fmt.Scan(&userInput)
	if scanError != nil {
		fmt.Printf("An error has ocurred, please try again\n")
		return
	}
	floatingPointNumber, conversionError := strconv.ParseFloat(userInput, 32)
	if conversionError != nil {
		fmt.Printf("This is not a valid input, please try again.\n")
		return
	}
	integerNumber := int(floatingPointNumber)
	fmt.Printf("Integer number: " + strconv.Itoa(integerNumber)+"\n")
}