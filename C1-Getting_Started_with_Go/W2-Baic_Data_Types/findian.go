/*
Task:
	Write a program which prompts the user to enter a string. The
	program searches through the entered string for the characters ‘i’, ‘a’, 
	and ‘n’. The program should print “Found!” if the entered string starts 
	with the character ‘i’, ends with the character ‘n’, and contains the
	character ‘a’. The program should print “Not Found!” otherwise. The
	program should not be case-sensitive, so it does not matter if the
	characters are upper-case or lower-case.

Examples:
	The program should print “Found!” for the following 
	example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. 
	The program should print “Not Found!” for the following strings, 
	“ihhhhhn”, “ina”, “xian”.
*/

package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	var input_string string

	fmt.Println("Enter a string:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input_string = scanner.Text()
	len_string := len(input_string)
	
	if len_string < 3 {
		fmt.Println("Not Found!")
		return
	} else {
		lower_string := strings.ToLower(input_string)
		if lower_string[0] == 105 &&
			lower_string[(len_string-1)] == 110 {
			if strings.Contains(lower_string, "a") {
				fmt.Println("Found!")
				return
			}
		}
		fmt.Println("Not Found!")
	}
}
