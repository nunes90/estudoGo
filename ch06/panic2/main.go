package main

import (
	"errors"
	"fmt"
)

func main() {
	msg := "good-bye"
	message(msg)
	fmt.Println("This line will not get printed")
}

func message(msg string) {
	if msg == "good-bye" {
		panic(errors.New("something went wrong"))
	}
}
