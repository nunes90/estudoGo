package main

import (
	"log"
)

// Here is a list of options for logging provided by the Go package that we can set in the function (https://go.dev/src/log/log.go?s=8483:8506#L28)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	name := "Thanos"
	log.Println("Demo app")
	log.Printf("%s is here!", name)
	log.Print("Run")
}
