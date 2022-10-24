/*
	Write a program which reads information from a file and represents 
	it in a slice of structs. Assume that there is a text file which contains a
	series of names. Each line of the text file has a first name and a last
	name, in that order, separated by a single space on the line. 

	Your program will define a name struct which has two fields, fname 
	for the first name, and lname for the last name. Each field will be a 
	string of size 20 (characters).

	Your program should prompt the user for the name of the text file. 
	Your program will successively read each line of the text file and 
	create a struct which contains the first and last names found in the
	file. Each struct created will be added to a slice, and after all lines
	have been read from the file, your program will have a slice
	containing one struct for each line in the file. After reading all lines 
	from the file, your program should iterate through your slice of 
	structs and print the first and last names found in each struct.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type Name struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
}

const (
	maxLength = 20
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func IterateSliceAndPrint(persons []Name) {
	for _, name := range persons {
		fmt.Printf("First name: %v, Last name: %v\n", name.Fname, name.Lname)
	}
}

func GetFileName(f *string) {
	fmt.Printf("Enter file name: ")
	l, err := fmt.Scanf("%s", f)
	if l == 0 {
		panic("ERROR: filename was not provided")
	}
	Check(err)
}

func ReadFileAndFillSplice(f string, persons *[]Name) {
	// Try to read file
	dat, err := os.ReadFile(f)
	Check(err)

	// Line to line scan
	scanner := bufio.NewScanner(strings.NewReader(string(dat)))
	Check(err)
	if scanner.Scan() {
		for scanner.Scan() {
			line := strings.Split(scanner.Text(), " ")
			if utf8.RuneCountInString(line[0]) <= maxLength &&
				utf8.RuneCountInString(line[1]) <= maxLength {
				var name Name
				name.Fname = line[0]
				name.Lname = line[1]
				*persons = append(*persons, name)
			} else {
				fmt.Println("fname or lname > 20 symbols: %v => skipped", line)
			}
		}
	} else {
		panic("ERROR with scanner")
	}
}

func main() {

	var persons []Name
	var f string

	GetFileName(&f)
	ReadFileAndFillSplice(f, &persons)
	IterateSliceAndPrint(persons)
}
