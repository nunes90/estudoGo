// Ex 1.05 - implementing a short variable declaration
package main

import (
	"fmt"
	"time"
)

func main() {
	Debug := false
	LogLevel := "info"
	startUpTime := time.Now()
	fmt.Println(Debug, LogLevel, startUpTime)
}
