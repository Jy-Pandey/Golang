package main

import (
	"encoding/json"
	"fmt"
)

// Step 1: Define the struct
type Student struct {
	Name  string `json:"name"` // treat Name as name while making json
	Password string `json:"-"`
	Email string `json:"email"`
	Age   int    `json:"age"`
	Courses []string `json:"enrolledCourses,omitempty"`
}
type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Lets learn how to encode Json data")
	// EncodeJson()
	DecodeJson()
}
func EncodeJson() {
	// Step 2: Put data into a slice
	students := []Student{
		{"jp", "abc12", "jp@gmail.com", 21, []string{"py", "js"}},
		{"ks", "abc12", "ks@gmail.com", 21, nil},
	}

	// Another way
	// p1 := Person{
	// 	Name:  "Jyoti",
	// 	Email: "jyoti@example.com",
	// 	Age:   22,
	// }

	// p2 := Person{
	// 	Name:  "Ankita",
	// 	Email: "ankita@example.com",
	// 	Age:   23,
	// }

	// people := []Person{p1, p2}

	// Step 3: Marshal to JSON
	res, err := json.MarshalIndent(students, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res)) // conver data from byte to string form
}
func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course // You are preparing a struct variable where you will store the decoded data.

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse) // This decodes the JSON and fills the data into the lcoCourse struct.
		fmt.Printf("%#v\n", lcoCourse) // # - Prints the complete struct with field names
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{} // This map will store keys as strings (like "coursename", "Price") and values as any type (interface{} = "any type" in Go).
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	// for k, v := range myOnlineData {
	// 	fmt.Printf("Key is %v  and value is %v and Type is: %T\n", k, v, v)
	// }

}