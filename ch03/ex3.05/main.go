// ex3.05 - Safely looping over a string
package main

import "fmt"

func main() {
	logLevel := "デバッグ"

	for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal))
	}

	username := "Sir_King_Über"
	// Length of a string
	fmt.Println("Bytes:", len(username))
	fmt.Println("Runes:", len([]rune(username)))

	// limit to 10 characters
	fmt.Println(string(username[:10]))
	fmt.Println(string([]rune(username)[:10]))
}
