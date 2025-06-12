package main

import (
	"fmt"
	"os"
	"strings"
)

func readingFile(filename string) {
	data, err := os.ReadFile(filename) //* take the path from current working directory
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Reading file:")
	fmt.Println(string(data))
}
func checkNilError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	fmt.Println("Welcome to file handling")

	filename := "temp.txt"

	//*Step:1 create and write into file(Overwrites if exists)
	//? Ye relative path hota hai — iska matlab:
	// "Is file ko current working directory (CWD) me dhundo ya yahi create karo"
	err := os.WriteFile(filename, []byte("Hello , Go Golangg\n"), 0644)
	checkNilError(err)

	//* Step2: Read from file
	readingFile(filename)

	//* Step3 : Append data to file first open it
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	checkNilError(err)

	file.WriteString("New line is Appended")
	defer file.Close() //! execute in end
	fmt.Println("Step 3: Data appended.")

	readingFile(filename)

	//*Step : Update file (replace text)
	content, _ := os.ReadFile(filename)
	updated := strings.Replace(string(content), "Go ", "Golang ", -1)
	os.WriteFile(filename, []byte(updated), 0644) // Overwrite

	//*Step 6: Delete the file
	// os.Remove(filename)
	// fmt.Println("Step 5: File deleted.")

}
