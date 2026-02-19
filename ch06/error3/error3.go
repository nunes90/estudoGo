package main

import (
	"errors"
	"fmt"
)

func main() {
	ErrBadData := errors.New("Some bad data")
	fmt.Printf("ErrBadData type: %T", ErrBadData)
}
