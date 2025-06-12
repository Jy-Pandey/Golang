package main

import "fmt"

func main() {
	fmt.Println("Learn Pointers")

	var ptr *int
	fmt.Println("Default Value of pointer is : ", ptr)
	// fmt.Println("Value of pointer is : ", *ptr)

	num := 23
	var ptr2 = &num
	fmt.Println("Value of pointer2 is : ", *ptr2)

	*ptr2 = *ptr2 + 2
	fmt.Println("My num is : ", num)

}
