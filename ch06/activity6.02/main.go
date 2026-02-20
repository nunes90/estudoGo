// activity6.02 - validating a bank customer's direct deposit submission
package main

import (
	"errors"
	"fmt"
	"strings"
)

type directDeposit struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

var (
	ErrInvalidLastName   = errors.New("invalid last name")
	ErrInvalidRoutingNum = errors.New("invalid routing number")
)

func main() {
	d := directDeposit{
		lastName:      "  ",
		firstName:     "Abe",
		bankName:      "XYZ Inc",
		routingNumber: 17,
		accountNumber: 1809,
	}

	err := d.validateRoutingNumber()
	if err != nil {
		fmt.Println(err)
	}
	err = d.validateLastName()
	if err != nil {
		fmt.Println(err)
	}
	d.report()

}

func (d *directDeposit) validateRoutingNumber() error {
	if d.routingNumber < 100 {
		return ErrInvalidRoutingNum
	}
	return nil
}

func (d *directDeposit) validateLastName() error {
	d.lastName = strings.TrimSpace(d.lastName)
	if len(d.lastName) == 0 {
		return ErrInvalidLastName
	}
	return nil
}

func (d *directDeposit) report() {
	fmt.Println(strings.Repeat("*", 80))
	fmt.Printf("Direct Deposit Details:\n")
	fmt.Printf("Last Name: %s\n", d.lastName)
	fmt.Printf("First Name: %s\n", d.firstName)
	fmt.Printf("Bank Name: %s\n", d.bankName)
	fmt.Printf("Routing Number: %d\n", d.routingNumber)
	fmt.Printf("Account Number: %d\n", d.accountNumber)
}
