// ex2.07 - expressionless switch statements
package main

import (
	"fmt"
	"time"
)

func main() {
	switch dayBorn := time.Sunday; {
	case dayBorn == time.Sunday || dayBorn == time.Saturday:
		fmt.Println("Born on the weekend")
	case 2+2 == 4:
		fmt.Println("math works")
	default:
		fmt.Println("Born some other day")
	}
}
