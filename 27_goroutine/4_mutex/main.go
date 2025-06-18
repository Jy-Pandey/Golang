package main

import (
	"fmt"
	"sync"
)

var counter = 0
var mu sync.Mutex // 🔒 create mutex

func increment(wg *sync.WaitGroup) {

	for i := 0; i < 1000; i++ {
		mu.Lock()   // 🚪 Lock the room
		counter++   // 👷‍♂️ Safe increment
		mu.Unlock() // 🔓 Unlock the room
	}
	wg.Done()
}
func RaceCond(wg *sync.WaitGroup) {

	for i := 0; i < 1000; i++ {
		counter++   // race condition
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg) //* 10 * 1000 = 10,000
	}

	wg.Wait()
	fmt.Println("✅ Final counter:", counter)
}
