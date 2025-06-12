package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case")

	rand.Seed(time.Now().UnixNano())

	num := rand.Intn(6) + 1 // generates number between 1 and 6
	fmt.Println("Random no. between 1-6 : ", num)
	// fmt.Println("Random int between 0-9:", rand.Intn(10))
	// fmt.Println("Random int:", rand.Int())
	// fmt.Println("Random float:", rand.Float64())
	// fmt.Println("Shuffled list of 5:", rand.Perm(5))

	// No need for break (Go automatically breaks after each case).
	// In Go, each case in a switch does NOT fall through by default — meaning, it only runs the matching case, then exits the switch.
	// But if you want to force the program to continue to the next case, even if its condition is false, you use fallthrough.

	// switch num {
	// case 1:
	// 	fmt.Println("Move one spot")
	// 	fallthrough
	// case 2:
	// 	fmt.Println("Move two spot")
	// case 3:
	// 	fmt.Println("Move three spot")
	// case 4:
	// 	fmt.Println("Move four spot")
	// case 5:
	// 	fmt.Println("Move five spot")
	// case 6:
	// 	fmt.Println("Move six spot, And again roll dice!")
	// 	fallthrough
	// default:
	// 	fmt.Println("What was that!")
	// }

	num2 := 5

	switch {
	case num2%2 == 0:
		fmt.Println("You rolled an even num2ber!")
		fallthrough
	case num2 == 1:
		fmt.Println("Move one spot")
		fallthrough
	case num2 == 2:
		fmt.Println("Move two spot")
	case num2 == 3:
		fmt.Println("Move three spot")
	case num2 == 4:
		fmt.Println("Move four spot")
	case num2 == 5:
		fmt.Println("Move five spot")
	case num2 == 6:
		fmt.Println("Move six spot, And again roll dice!")
		fallthrough
	default:
		fmt.Println("What was that!")
	}
}
