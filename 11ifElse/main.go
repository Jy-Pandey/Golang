package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}

	if num := 0; num%2 == 0 {
		fmt.Println("EVen num")
	} else {
		fmt.Println("Odd num")
	}
}
