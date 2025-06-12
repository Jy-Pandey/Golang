package main

import "fmt"

func main() {
	fmt.Println("LEarn maps")

	// way 1 - without make()
	lang := map[string]string{"js": "Javascript", "py": "python"}
	fmt.Println("language list is : ", lang)

	// way2 - using make
	color := make(map[int]string)
	color[1] = "Blue"
	color[2] = "Red"
	color[3] = "Green"
	fmt.Printf("Colors list is %v\n", color)

	// Accessing values
	fmt.Println("First color is:", color[1])

	// Adding a new key-value
	color[4] = "Yellow"

	// Updating a value
	color[2] = "Orange"

	// Delete a color
	delete(color, 3) // delete green color

	// Iterate to map
	for key, val := range color {
		fmt.Printf("Key : %v, Val : %v\n", key, val)
	}
	// fmt.Printf("Colors list is %v", color)

}
