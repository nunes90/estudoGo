// activity4.04 - Removing an element from a slice
package main

import "fmt"

func removeBad() []string {
	sli := []string{"Good", "Good", "Bad", "Good", "Good"}
	sli = append(sli[:2], sli[3:]...)
	return sli
}

func main() {
	fmt.Println(removeBad())
}
