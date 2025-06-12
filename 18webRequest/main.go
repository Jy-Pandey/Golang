package main

import (
	"fmt"
	"io"
	"net/http"
)

const url string = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	fmt.Println("Welcome to web Request in Golang")

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close() // Developer need to manually close the connection

	dataByte, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}
	content := string(dataByte)
	fmt.Println("Response : ", content)
}
