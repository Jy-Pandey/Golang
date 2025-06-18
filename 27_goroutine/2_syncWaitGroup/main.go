package main

import (
	"fmt"
	"sync"
)
func worker(w int, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Printf("Worker %d started\n", w)
	// perform any work
	fmt.Printf("Worker %d eneded\n", w)
}
func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // add 1 go routine to list
		go worker(i, &wg)
	}

	wg.Wait() // wait for all goroutine to complete
	fmt.Println("Main fn ended")
}
