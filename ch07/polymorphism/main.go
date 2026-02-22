// You can achieve polymorphism in Go by using interfaces.
// by being able to pass various concrete types to a method or function that
// has an interface type as an input parameter.

// Continue part 2 on polymorphism2/main.go

package main

import "fmt"

type Speaker interface {
	Speak() string
}

type cat struct{}

func main() {
	c := cat{}
	catSpeak(c)
}

func (c cat) Speak() string {
	return "Purr Meow"
}

func catSpeak(c cat) { // tipo concreto, nao interface
	fmt.Println(c.Speak())
}
