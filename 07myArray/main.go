package main

import "fmt"

func main() {
	fmt.Println("My Arrays")

	var fruit = [4]string{"App", "Man", "Ban"}
	// var fruit [4]string
	// fruit[0] = "App"
	// fruit[2] = "Mang"

	// fmt.Println("Enter values ")
	// for i := 0; i < 4; i++ {
	// 	fmt.Scanln(&fruit[i])
	// }
	// for i := range fruit {
	// 	fmt.Scanln(&fruit[i])
	// }

	fmt.Println("Fruit lists are : ", fruit)
	fmt.Println("Fruit list length is : ", len(fruit))

}
