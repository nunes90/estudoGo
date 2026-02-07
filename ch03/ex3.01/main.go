// ex3.01 - Program to measure password complexity
package main

import (
	"fmt"
	"unicode"
)

/*
requirements:
- Have a lowercase letter
- Have an uppercase letter
- Have a number
- Have a symbol
- Be 8 or more characters (15)
*/

func passwordChecker(pw string) bool {
	pwR := []rune(pw)

	if len(pwR) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, v := range pwR {
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	if passwordChecker("") {
		fmt.Println("password good")
	} else {
		fmt.Println("password bad")
	}

	if passwordChecker("This!I5A") {
		fmt.Println("password good")
	} else {
		fmt.Println("password bad")
	}

}
