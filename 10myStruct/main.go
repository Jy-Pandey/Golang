package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade string
}

func main() {
	fmt.Println("Learn Struct")

	s1 := Student{Name: "Alice", Age: 21, Grade: "A"}
	fmt.Printf("Student details %v\n", s1)
	fmt.Printf("Student complete details %+v\n", s1)

	// Accessing fields
	fmt.Println("Name:", s1.Name)
	fmt.Println("Age:", s1.Age)
	fmt.Println("Grade:", s1.Grade)

	// Updating fields
	s1.Age = 22
	fmt.Println("Updated Age:", s1.Age)

	// Map of struct
	studentMap := map[string]Student{
		"Alice": {Name: "Alice", Age: 21, Grade: "A"},
		"Bob":   {Name: "Bob", Age: 19, Grade: "C"},
	}
	fmt.Println("Student details are : ", studentMap)
	fmt.Printf("Alice details are : %+v \n", studentMap["Alice"])
	fmt.Printf("Alice Age is : %v \n", studentMap["Alice"].Age)
}
