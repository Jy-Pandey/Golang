package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {
	// 	fmt.Printf("index is  and value is %v\n", day)
	// }

	idx := 1

	for idx < 10 {

		if idx == 10 {
			goto lco
		}

		if idx == 5 {
			idx++
			continue
		}

		fmt.Println("Value is: ", idx)
		idx++
	}

lco:
	fmt.Println("Jumping at label lco")

}
