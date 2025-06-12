package main

import "fmt"

func main() {
	fmt.Println("Learn Defer in golang")

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i) // stack  - 0 1 2 3 4 => out 43210
	}
}
