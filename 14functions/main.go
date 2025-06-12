package main

import "fmt"

func divideAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}

// if all arguments of same type u can write only once in last
func multiply(a, b int) (result int) {
	result = a * b
	return // no need to explicitly write `return result`
}
func sumAll(values ...int) int {

	tot := 0
	for _, val := range values {
		tot += val
	}
	return tot
}
func main() {
	quotient, remainder := divideAndRemainder(10, 3)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)

	fmt.Println("Muttiplication of 2 * 4 is : ", multiply(2, 4))

	// arr := []int{1, 2, 3, 4}
	// variadic function
	tot := sumAll(1, 2, 3, 4)
	fmt.Println("Sum is : ", tot)

	// Anonymous(UnNamed) function
	greet := func(name string) {
		fmt.Println("Hello", name)
	}

	greet("Go developer!")

	sum := func(x, y int) int { return x + y }
	fmt.Println("Result:", applyOperation(3, 4, sum)) // Output: Result: 7

}
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}
