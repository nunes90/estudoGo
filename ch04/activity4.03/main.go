// activity4.03 - Slicing the week
package main

import "fmt"

var week = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

func main() {
	week = append(week[6:], week[:6]...)
	fmt.Println(week)
}
