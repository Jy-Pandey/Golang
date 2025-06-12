package main

import "fmt"

type Rectangle struct {
	length int
	bredth int
}

func (r Rectangle) area() int {
	return r.length * r.bredth
}

func (r *Rectangle) scale(factor int) {
	r.length = factor * r.length
	r.bredth = factor * r.bredth
}

func main() {
	r := Rectangle{3, 4}
	fmt.Println("Area:", r.area()) // 12

	r.scale(2)
	fmt.Println("Scaled Area:", r.area()) // 48
}
