package main

import (
	"log"
	"errors"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("Start of our app")
	err := errors.New("Application Aborted!")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("End of our app")
}
