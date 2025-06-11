package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Learn Time")

	now := time.Now()
	fmt.Println("now : ", now.AddDate(2, 0, 0))
	fmt.Println("Year is : ", now.Year())
	fmt.Println("Month is : ", now.Month())
	fmt.Println("Hour is : ", now.Hour())
	fmt.Println("Formatted Date : ", now.Format("01-02-2006 Monday 15:04:05"))

	fmt.Println(time.Hour)
	fmt.Println(2 * time.Hour)
	future := now.Add(2 * time.Hour)
	fmt.Println("Future time after two hourse : ", future)

	// sleep
	// fmt.Println("Waiting for 5 seconds..")
	// time.Sleep(5 * time.Second)
	// fmt.Println("Done!")

	input := "11-06-2025"
	parsedTime, err := time.Parse("01-02-2006", input)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed time:", parsedTime)
	}

	// Difference between two time
	t1 := time.Now()
	time.Sleep(1 * time.Second)
	t2 := time.Now()

	diff := t2.Sub(t1)
	fmt.Println("Time difference:", diff)

}
