package main

import (
	"fmt"
	"time"
)

var (
	Debug       bool         // no value, default false
	LogLevel    = "info"     // no type, initialized
	startUpTime = time.Now() // no type, initialized
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)

}
