package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	maxInputLength = 10
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadStdinLineToStringArray(stringArray *[]string) {
	fmt.Printf("Input sequence up to 10 integers separated by space: ")
	scanner := bufio.NewReader(os.Stdin)
	inputString, err := scanner.ReadString('\n')
	Check(err)
	if len(inputString) == 0 {
		fmt.Println("Empty input, repeat")
		os.Exit(1)
	}
	inputString = strings.Trim(inputString, "\n")
	*stringArray = strings.Split(inputString, " ")
	if len(*stringArray) > 10 {
		fmt.Println("Repeat with up to 10 integers")
		os.Exit(1)
	}
}

func ConvertStringArrayToIntegerArray(stringArray *[]string, integerArray *[]int) {
	for _, char := range *stringArray {
		i, err := strconv.Atoi(char)
		Check(err)
		*integerArray = append(*integerArray, i)
	}
}

func Swap(integerArray *[]int, i int){
	(*integerArray)[i], (*integerArray)[i+1] = (*integerArray)[i+1], (*integerArray)[i]
}

func BubbleSort(integerArray *[]int) {
	var sorted bool
	for i := 0; i < len(*integerArray); i++ {
		sorted = true
		for j := 0; j < len(*integerArray)-1-i; j++ {
			if (*integerArray)[j] > (*integerArray)[j+1] {
				Swap(integerArray, j)
				fmt.Printf("BubbleSort iteration (%v), swap (%v), sort result: %v\n", i, j, *integerArray)
				sorted = false
			}
		}
		if sorted{
			fmt.Printf("BubbleSort result achieved on: iteration (%v), sort result: %v\n", i-1, *integerArray)
			break
		}
	}
}

func main() {

	var stringArray []string
	var integerArray []int

	ReadStdinLineToStringArray(&stringArray)
	ConvertStringArrayToIntegerArray(&stringArray, &integerArray)
	BubbleSort(&integerArray)
}
