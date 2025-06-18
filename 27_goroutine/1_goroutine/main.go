package main

import (
	"fmt"
	"time"
)

func fun1()  {
	time.Sleep(1 * time.Second)
	fmt.Println("Func1 : First line :)")
}

func fun2()  {
	time.Sleep(2 * time.Second)
	fmt.Println("Func2 : Hello world!")

}
func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println("Number:", i)
        time.Sleep(time.Second)
    }
}
func main() {
	fmt.Println("Learning goroutine")
	go fun2()
	printNumbers() // after end of this function control goes to next line
	go fun1()

	time.Sleep(1200* time.Millisecond); // wwait for 2 sec
}
