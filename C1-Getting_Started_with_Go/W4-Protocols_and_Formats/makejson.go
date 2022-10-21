/*
	Write a program which prompts the user to first enter a name, and 
	then enter an address. Your program should create a map and add
	the name and address to the map using the keys “name” and 
	“address”, respectively. Your program should use Marshal() to create
	a JSON object from the map, and then your program should print the 
	JSON object.
*/

package main

import (
	"os"
	"fmt"
	"bufio"
	"encoding/json"
	
)

func scan_input(user_message string) (string){
	fmt.Println(user_message)	
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	} else {
		panic("ERROR with scanner")
	}
	return input
}

func create_json(m map[string]string) ([]byte){
	u, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return u
}

func main() {

	var m map[string]string

	m = make(map[string]string)
	m["name"] = scan_input("Enter name : ")
	m["address"] = scan_input("Enter address : ")

	json := create_json(m)

	fmt.Println(json)
	fmt.Println(string(json))
}
