// Ex 1.06 - declaring multiple variables from a function
package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	Debug, LogLevel, startUpTime := getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)
}
