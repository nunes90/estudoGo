package main

import (
	"fmt"
	"time"
)

func main() {
	timeToManipulate := time.Now()
	toBeAdded := time.Duration(10 * time.Second)
	fmt.Println("The original time:", timeToManipulate)
	fmt.Printf("%v duration later %v", toBeAdded, timeToManipulate.Add(toBeAdded))
}
