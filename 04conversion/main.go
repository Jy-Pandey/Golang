package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Type conversion")

	fmt.Println("Enter your Rating for pizza : ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for your Rating,", input)
	// Type of input is string

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your rating is incremented by 1, ", numRating+1)
	}

	// string to int
	var s string = "23"
	i, err := strconv.Atoi(s)

	if err != nil {
		fmt.Println("Error while converting string to int : ", err)
	} else {
		fmt.Println("Int value+1 is :", i+1)
	}

}
