package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Learn about User Input")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for pizza :")

	// 	os.Stdin → Refers to the standard input (keyboard).
	// bufio.NewReader(...) → Creates a buffered reader, which can read data more flexibly. => allow spaces

	// Ok comma // Ok err syntax
	// input and err
	// Reads the whole line, including spaces, until Enter is pressed
	input, _ := reader.ReadString('\n') // read untill newline
	fmt.Println("Thanks for your rating: ", input)
	fmt.Printf("Type of input is %T \n", input)

	var name string
	fmt.Println("Enter your name :")
	// read untill a space
	fmt.Scanln(&name)
	fmt.Printf("Welcome %s!", name)
}
