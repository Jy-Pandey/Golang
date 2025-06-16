package main

import "fmt"

//? ////////////////////////////////////////////////////////
// STEP 1: Define the Interface
//? ////////////////////////////////////////////////////////

// Payment is an interface.
// It says: "Any type that has a Pay(float64) method is a Payment."
type Payment interface {
	Pay(amount float64)
}

//? ////////////////////////////////////////////////////////
// STEP 2: Create Types That Implement the Interface
//? ////////////////////////////////////////////////////////

// CreditCard is a struct representing credit card payment.
type CreditCard struct{}

// Pay is the method of CreditCard that matches the Payment interface.
func (c CreditCard) Pay(amount float64) {
	fmt.Println("Paid with credit card:", amount)
}

// PayPal is another struct representing PayPal payment.
type PayPal struct{}

// Pay is the method of PayPal that also matches the Payment interface.
func (p PayPal) Pay(amount float64) {
	fmt.Println("Paid with PayPal:", amount)
}

//? ////////////////////////////////////////////////////////
// STEP 3: Function That Uses the Interface
//? ////////////////////////////////////////////////////////

// ProcessPayment accepts anything that satisfies the Payment interface.
// It doesn't care *how* you pay — just that you *can* pay.
func ProcessPayment(p Payment) {
	p.Pay(100) // Process a payment of 100
}

//? ////////////////////////////////////////////////////////
// STEP 4: Main Function
//? ////////////////////////////////////////////////////////

func main() {
	// Create a CreditCard object
	cc := CreditCard{}

	// Create a PayPal object
	pp := PayPal{}

	// Process payment using credit card
	ProcessPayment(cc) // Output: Paid with credit card: 100

	// Process payment using PayPal
	ProcessPayment(pp) // Output: Paid with PayPal: 100
}
