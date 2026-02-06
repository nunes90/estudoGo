// Ex 1.03 - declaring multiple variables at once with var
package main

import (
	"fmt"
	"time"
)

var (
	Debug       bool      = false
	LogLevel    string    = "info"
	startUpTime time.Time = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}
