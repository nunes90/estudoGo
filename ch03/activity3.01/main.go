// activity3.01 - Sales tax calculator
package main

import "fmt"

/*
Item   | Cost  | Sales Tax Rate
Cake   | $0.99 | 7.5%
Milk   | $2.75 | 1.5%
Butter | $0.87 | 2%
*/

func main() {
	taxTotal := .0
	// Cake
	taxTotal += salesTax(.99, .075)
	// Milk
	taxTotal += salesTax(2.75, .015)
	// Butter
	taxTotal += salesTax(.87, .02)
	// Total
	fmt.Println("Sales Tax Total: ", taxTotal)

}

func salesTax(cost float64, taxRate float64) float64 {
	return cost * taxRate
}
