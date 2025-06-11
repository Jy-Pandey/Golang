package main

import (
	"fmt"
	"mylearning/myutil"
)

func main() {
	fmt.Println("Hello, World!")

	name := "Jyoti"
	age := 20

	// fmt.Print("Hello!") // In a single line
	fmt.Printf("Hello %s, Your age is %d", name, age)
	fmt.Println()
	myutil.PrintMessage("Called from main.go")

	// array
	arr := []int{1, 2, 3, 4}

	for i, v := range arr {
		fmt.Println(i, v)
	}

}
