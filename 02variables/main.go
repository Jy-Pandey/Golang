package main

import "fmt"

// Allowed
var num = 2
var num2 int = 5

// Not Allowed
// num3 := 5

func main() {
	fmt.Println("variables")
	var username string = "Jyoti"
	fmt.Printf("Type of variable is : %T \n", username)

	var smallVal uint8 = 255 // 0 to 255
	fmt.Printf("Type of variable is : %T \n", smallVal)

	var smallFloat float32 = 199.43215657 // upto 5 precison only
	fmt.Println(smallFloat)
	fmt.Printf("Type of variable is : %T \n", smallFloat)

	// Default values and alisese
	var anotherVar int
	fmt.Println("Default Val : ", anotherVar)
	fmt.Printf("Type of variable is : %T \n", anotherVar)

	var defFloat float32
	fmt.Println("Default Val : ", defFloat)
	fmt.Printf("Type of variable is : %T \n", defFloat)

	var defStr string
	fmt.Println("Default Val : ", defStr)
	fmt.Printf("Type of variable is : %T \n", defStr)

	// implicit type
	var website = "AppPerfect.com"
	fmt.Println(website)
	// website = 5 // Wrong

	// This syntax is only allowed inside a fuction or method
	marks := 80
	fmt.Println("Marks is :", marks)
}
