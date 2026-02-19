package main

import (
	"errors"
	"fmt"
)

func main() {
	// try access from errors package
	es := errors.errorString{}
	es.s = "slacker"
	fmt.Println(es)
}
