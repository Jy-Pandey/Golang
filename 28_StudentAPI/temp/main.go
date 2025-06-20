package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// Set expiration time to 1 hour from now
	expirationTime := time.Now().Add(1 * time.Hour)

	// Convert to NumericDate
	numericDate := jwt.NewNumericDate(expirationTime)
	// NumericDate: 2025-06-20 18:25:45 +0530 IST


	fmt.Println("NumericDate:", numericDate)

	// Check if the date is valid (not expired)
	if numericDate.Time.Before(time.Now()) {
		fmt.Println("Token has expired!")
	} else {
		fmt.Println("Token is still valid!")
	}
}
