package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func PerformGetReq() {
	const myurl string = "http://localhost:8000/get"

	res, err := http.Get(myurl)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response status : ", res.Status)

	dataByte, _ := io.ReadAll(res.Body)
	fmt.Println("WAY 1 : Data is - ", string(dataByte))

	// WAY 2 - using strings.builder - If you're combining multiple strings or byte slices, strings.Builder is better.
	// You’re adding in parts or inside loops then use strings.builder otherwise string(data) is good

	var resString strings.Builder
	byteCnt, _ := resString.Write(dataByte)
	fmt.Println("Byte cnt is ", byteCnt)
	fmt.Println("WAY 2 : Content is : ", resString.String())
}

func PerformPostJsonReq() {
	const myurl string = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"name" : "Jyoti",
			"age" : "21"
		}
	`)

	res, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	fmt.Println("Post req data is : ", string(content))
}
func PerformPostFormReq() {
	const myurl string = "http://localhost:8000/postform"

	formData := url.Values{} // reperesent data like this map[string][]string

	formData.Add("name", "jyoti") // add this data to map only.. allow duplicate keys
	formData.Set("age", "22") // stores only unique key values .. overrite previous one if exist
	formData.Set("emaill" , "jp@appperfect.com")

	res, err := http.PostForm(myurl, formData)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	content, _ := io.ReadAll(res.Body)
	fmt.Println("Post formData req data is : ", string(content))
}
func main() {
	fmt.Println("Http Requests")
	// PerformGetReq()
	// PerformPostJsonReq()
	PerformPostFormReq()
}
