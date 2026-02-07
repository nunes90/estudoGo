// activity2.01 - looping over map data using range
package main

import "fmt"

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	var countMaior int
	var wordMaior string
	for word, count := range words {
		if count > countMaior {
			countMaior = count
			wordMaior = word
		}
	}
	fmt.Println("Most popular word: ", wordMaior)
	fmt.Println("With a count of: ", countMaior)
}
