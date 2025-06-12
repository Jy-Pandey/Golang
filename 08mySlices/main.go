package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Let's learn Slices")

	var s1 []int
	s2 := []int{}
	fmt.Println(s1 == nil)
	fmt.Printf("Type : %T \n", s1)
	fmt.Println(s2 == nil)

	sl1 := make([]int, 3, 5)
	fmt.Println("Slice length is : ", len(sl1))
	fmt.Println("Slice capacity is : ", cap(sl1))

	sl2 := append(sl1, 20, 40)
	fmt.Println("updated slice length is : ", len(sl2))
	fmt.Println("updated slice capacity is : ", cap(sl2))

	// slice to still have reference to sl1
	sl2[1] = 100
	fmt.Println("sl1 is ", sl1)
	fmt.Println("sl2 is ", sl2)

	// This time we are appending values beyond the sl1 capacity, so capacity gets double = 5 * 2 = 10
	// So a new array will be created internally and sl3 will have a reference to it
	sl3 := append(sl1, 20, 40, 50)
	fmt.Println("updated slice3 length is : ", len(sl3))
	fmt.Println("updated slice3 capacity is : ", cap(sl3))

	sl3[1] = 400
	fmt.Println("sl1 is ", sl1)
	fmt.Println("sl2 is ", sl2)
	fmt.Println("sl3 is ", sl3)

	// sort the slice
	fmt.Println(sort.IntsAreSorted(sl3))
	sort.Ints(sl3)
	fmt.Println("sl3 is ", sl3)
	fmt.Println(sort.IntsAreSorted(sl3))

	// Deletion from slice
	// Delete the given index element
	idx := 3
	sl3 = append(sl3[:idx], sl3[idx+1:]...)
	fmt.Println("sl3 is after deletion is ", sl3)
}
