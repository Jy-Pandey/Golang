package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to urls in golang")

	res, _ := url.Parse(myurl)
	// fmt.Println(res.Scheme)
	// fmt.Println(res.Host)
	// fmt.Println(res.Port())
	// fmt.Println(res.Path)
	// fmt.Println(res.RawQuery)
	// fmt.Println(res.Query())

	qparams := res.Query() // gives the map
	fmt.Println("Course name : ", qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	//* create url from parts of url

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawQuery: "user=hiteh",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
