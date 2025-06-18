package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadFile(file string, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done() // run at the last
	time.Sleep(time.Second) // 1 sec delay
	res := fmt.Sprintf("Downloaded file %s", file) // return formatted string

	ch <- res // Send result into channel
}
func main() {

	files := []string{"file1.pdf", "file2.png", "file3.mp4"}
	ch := make(chan string)
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1) // add one goroutine
		go downloadFile(file, ch, &wg)
	}

	// Another goroutine
	go func(){ // anonymous function (IIFE - Immediately Invoked Function Expression) if we dont use this function channel will never close and next for loop cause infinite loop
		wg.Wait()
		close(ch)
	}()


	for msg := range ch {
		fmt.Println("channel res :", msg)
	}

	fmt.Println("Main function ended")
}