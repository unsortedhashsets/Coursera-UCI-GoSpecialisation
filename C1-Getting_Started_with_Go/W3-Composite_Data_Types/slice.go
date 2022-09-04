/*
	Write a program which prompts the user to enter integers and stores
	the integers in a sorted slice. The program should be written as a 
	loop. Before entering the loop, the program should create an empty 
	integer slice of size (length) 3. During each pass through the loop, the 
	program prompts the user to enter an integer to be added to the 
	slice. The program adds the integer to the slice, sorts the slice, and
	prints the contents of the slice in sorted order. The slice must grow in
	size to accommodate any number of integers which the user decides 
	to enter. The program should only quit (exiting the loop) when the 
	user enters the character ‘X’ instead of an integer.
*/

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func append_and_sort(input string, slice []int, sliceBool *[]bool, index_addr *int) []int {
	var done = false;
	input_int, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	for i, b := range *sliceBool {
		if b == false && !done {
			slice[i] = input_int
			(*sliceBool)[i] = true
			done = true
			fmt.Println(slice)
			fmt.Println(*sliceBool)
		}
	}
	
	if !done{
		slice = append(slice, input_int)
	}
	
	sort.Slice(slice, func(i, j int) bool {
		if slice[i] < slice[j]{
			if len(slice) == len(*sliceBool){
				(*sliceBool)[j], (*sliceBool)[i] = (*sliceBool)[i], (*sliceBool)[j]
				fmt.Println(slice)
				fmt.Println(*sliceBool)
			}
			return true
		} else {
			return false
		}
	})
	fmt.Println(slice)
	return slice
}

func check_and_convert(input_addr *string, integer_addr *int) bool {
	if *input_addr == "X" {
		os.Exit(0)
	}
	if value, err := strconv.Atoi(*input_addr); err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Printf("Input integer is: %d\n", value)
		*integer_addr = value
		return true
	}
}

func scan_input(input_addr *string) bool {

	fmt.Printf("Type an integer number:\n")
	len, err := fmt.Scanf("%v", input_addr)

	if len == 0 {
		fmt.Printf("Input is empty, please try again\n")
		return false
	} else {
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {

			return true
		}
	}
	return false
}

func main() {
	var slice = make([]int, 3, 3)
	var sliceBool = make([]bool, 3, 3) 
	var input string
	var integer int
	var index int = 0

	for {
		// Scan input
		if scan_input(&input) {
			// check and convert to int
			if check_and_convert(&input, &integer) {
				// Append and sort
				slice = append_and_sort(input, slice, &sliceBool, &index)
			}
		}
	}
}